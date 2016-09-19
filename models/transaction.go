package models

import (
	"time"

	"bitbucket.org/SummerCampDev/summercamp/models/utils"
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
	ID         int             `json:"id" orm:"column(id)"`
	UserID     int             `json:"user_id" orm:"column(user_id)"`
	ProjectID  int             `json:"project_id" orm:"column(project_id)"`
	Type       TransactionType `json:"type" orm:"column(type)"`
	Amount     int             `json:"amount" orm:"column(amount)"`
	Comment    string          `json:"comment" orm:"column(comment)"`
	CreateTime time.Time       `json:"create_time" orm:"column(create_time);auto_now_add;type(datetime)"`
}

func (t *Transaction) Save() bool {
	var err error
	var action string

	if t.ID == 0 {
		_, err = DB.Insert(t)
		action = "create"
	} else {
		_, err = DB.Update(t)
		action = "update"
	}

	return utils.ProcessError(err, action+" transaction")
}

// Delete deletes the transaction record from the db
func (t *Transaction) Delete() bool {
	_, err := DB.Delete(t)
	return utils.ProcessError(err, "delete user")
}
