package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type TransactionModel_20160919_230037 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &TransactionModel_20160919_230037{}
	m.Created = "20160919_230037"
	migration.Register("TransactionModel_20160919_230037", m)
}

// Run the migrations
func (m *TransactionModel_20160919_230037) Up() {
	m.SQL(`CREATE TYPE transaction_type AS ENUM ('balance', 'transaction');
		CREATE TABLE public.transactions
(
    id SERIAL PRIMARY KEY NOT NULL,
    user_id INT NOT NULL,
    project_id INT,
    type TRANSACTION_TYPE NOT NULL,
    amount INT CHECK (amount<>0),
    comment INT,
    create_time TIMESTAMP DEFAULT now() NOT NULL,
    CONSTRAINT transactions_users_id_fk FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE UNIQUE INDEX transactions_id_uindex ON public.transactions (id);
COMMENT ON COLUMN public.transactions.amount IS 'the value here is in cents, so when we want to know the dollar value, we divide this field by 100
';
`)

}

// Reverse the migrations
func (m *TransactionModel_20160919_230037) Down() {
	m.SQL(`DROP TYPE IF EXISTS transaction_type;
	DROP TABLE IF EXISTS public.transactions;
	`)
}
