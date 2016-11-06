package models

import (
	"time"

	"bitbucket.org/SummerCampDev/summercamp/models/utils"
	"github.com/astaxie/beego"
)

type TransactionType string

const (
	TransactionTypeBalance  TransactionType = "balance"
	TransactionTypeTransfer TransactionType = "transfer"
)

func (tt TransactionType) Valid() bool {
	return tt == TransactionTypeBalance || tt == TransactionTypeTransfer
}

type Transaction struct {
	ID          int             `json:"id" orm:"column(id)"`
	UserID      int             `json:"user_id" orm:"column(user_id)"`
	ProjectID   int             `json:"project_id" orm:"column(project_id)"`
	Type        TransactionType `json:"type" orm:"column(type)"`
	Amount      int             `json:"amount" orm:"column(amount)"`
	Description string          `json:"description" orm:"column(description)"`
	CreateTime  time.Time       `json:"create_time" orm:"column(create_time);auto_now_add;type(datetime)"`
}

func (t *Transaction) TableName() string {
	return "transactions"
}

// Save inserts a new or updates an existing transaction record in the DB.
func (t *Transaction) Save() bool {
	_, err := DB.InsertOrUpdate(t)
	return utils.ProcessError(err, "insert or update transaction")
}

// Delete deletes the transaction record from the db
func (t *Transaction) Delete() bool {
	_, err := DB.Delete(t)
	return utils.ProcessError(err, "delete user")
}

type TransactionsAPI struct{}

var Transactions *TransactionsAPI

func (t *TransactionsAPI) FetchByID(id int) (*Transaction, bool) {
	transaction := Transaction{ID: id}
	err := DB.Read(&transaction)
	return &transaction, utils.ProcessError(err, "fetch the transaction by id")
}

func (t *TransactionsAPI) NewBalance(userID, amount int, description string) *Transaction {
	return &Transaction{
		UserID:      userID,
		Amount:      amount,
		Type:        TransactionTypeBalance,
		Description: description,
	}
}

func (t *TransactionsAPI) NewTransfer(userID, projectID, amount int, description string) *Transaction {
	return &Transaction{
		UserID:      userID,
		ProjectID:   projectID,
		Amount:      amount,
		Type:        TransactionTypeTransfer,
		Description: description,
	}
}

func (t *TransactionsAPI) FetchTransferByUserID(userID int) ([]Transaction, bool) {
	return t.fetchTransactionsByUserIDAndType(userID, TransactionTypeTransfer)
}

func (t *TransactionsAPI) FetchBalanceByUserID(userID int) ([]Transaction, bool) {
	return t.fetchTransactionsByUserIDAndType(userID, TransactionTypeBalance)
}

func (t *TransactionsAPI) FetchDepositsByUserID(userID int) ([]Transaction, bool) {
	return t.fetchTransactionsByUserIDAndType(userID, TransactionTypeBalance, "amount>0")
}

func (t *TransactionsAPI) FetchWithdrawalsByUserID(userID int) ([]Transaction, bool) {
	return t.fetchTransactionsByUserIDAndType(userID, TransactionTypeBalance, "amount<0")
}

func (t *TransactionsAPI) fetchTransactionsByUserIDAndType(userID int, ttype TransactionType, extraFilters ...interface{}) ([]Transaction, bool) {
	if !ttype.Valid() {
		beego.BeeLogger.Error("invalid transaction type. Type: %v", ttype)
		return nil, false
	}
	query := "SELECT * FROM transactions WHERE user_id=$1 AND type=$2"

	if extraFilters != nil {
		for _, filterI := range extraFilters {
			if filter, ok := filterI.(string); ok {
				query += " AND " + filter
			}
		}
	}
	var transactions []Transaction

	query += " ORDER BY create_time DESC;"
	_, err := DB.Raw(query, userID, ttype).QueryRows(&transactions)

	return transactions, utils.ProcessError(err, " fetch transactions")
}
