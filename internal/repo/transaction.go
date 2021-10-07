package repo

import (
	"log"
	"time"

	"github.com/danisbagus/golang-caching-redis/internal/domain"
	"github.com/jmoiron/sqlx"
)

type ITransactionRepo interface {
	FetchAll() ([]domain.Transaction, error)
}

type TransactionRepo struct {
	db *sqlx.DB
}

func NewTransactionRepo(db *sqlx.DB) ITransactionRepo {
	return &TransactionRepo{
		db: db,
	}
}

func (r TransactionRepo) FetchAll() ([]domain.Transaction, error) {
	transactions := make([]domain.Transaction, 0)

	fetchAllTrasactionQuery := `select * from transactions`

	err := r.db.Select(&transactions, fetchAllTrasactionQuery)

	if err != nil {
		log.Printf("Error while quering find all purchase transaction by merchant id " + err.Error())
		return nil, err
	}

	time.Sleep(3 * time.Second)

	return transactions, nil
}
