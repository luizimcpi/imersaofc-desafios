package repository

import (
	"database/sql"
	"time"
)

type TransactionRepositoryDb struct {
	db *sql.DB
}

func NewTransactionRepositoryDb(db *sql.DB) *TransactionRepositoryDb {
	return &TransactionRepositoryDb{db: db}
}

func (t *TransactionRepositoryDb) Insert(id string, account string, amount float64) error {
	stmt, err := t.db.Prepare(`
		insert into transactions (id, account_id, amount, created_at)
		values($1,$2,$3,$4)
		`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		id,
		account,
		amount,
		time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
}
