package database

import (
	"testing"

	"bitbucket.org/SummerCampDev/summercamp/models"
	"bitbucket.org/SummerCampDev/summercamp/tests/setup"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTransactionsAPI(t *testing.T) {
	Convey("Test transactions API", t, func() {
		Convey("fetch transfer by user id", func() {
			transactions, ok := models.Transactions.FetchTransferByUserID(1)

			So(ok, ShouldBeTrue)
			So(transactions, ShouldHaveLength,
				setup.GetFixture("transactions").Filter("type", setup.Equal, "transfer").Count())
		})

		Convey("fetch balance by user id", func() {
			transactions, ok := models.Transactions.FetchBalanceByUserID(1)

			So(ok, ShouldBeTrue)
			So(transactions, ShouldHaveLength,
				setup.GetFixture("transactions").Filter("type", setup.Equal, "balance").Count())
		})

		Convey("fetch withdrawls by user id", func() {
			transactions, ok := models.Transactions.FetchWithdrawalsByUserID(1)

			So(ok, ShouldBeTrue)
			So(transactions, ShouldHaveLength, 1) //TODO: fixtures_lib for ints
		})

		Convey("fetch deposits by user id", func() {
			transactions, ok := models.Transactions.FetchDepositsByUserID(1)

			So(ok, ShouldBeTrue)
			So(transactions, ShouldHaveLength, 1) //TODO: fixtures_lib for ints
		})

		Convey("fetch all", func() {
			transactionID := 1
			transaction, ok := models.Transactions.FetchByID(transactionID)

			So(ok, ShouldBeTrue)
			So(transaction, ShouldNotBeNil)
			So(transaction.ID, ShouldEqual, transactionID)
		})
	})
}

func TestTransactionModel(t *testing.T) {
	Convey("Test transaction model", t, func() {
		transaction := models.Transaction{
			UserID:    1,
			ProjectID: 1,
			Type:      "balance",
			Amount:    100,
		}

		ok := transaction.Save()
		So(ok, ShouldBeTrue)
		So(transaction.ID, ShouldNotEqual, 0)

		ok = transaction.Delete()
		So(ok, ShouldBeTrue)

	})
}
