package usecase

import (
	"github.com/danisbagus/golang-caching-redis/internal/dto"
	"github.com/danisbagus/golang-caching-redis/internal/repo"
)

type ITransactionUsecase interface {
	GetAll() (*dto.TransactionListResponse, error)
}

type TransactionUsecase struct {
	repo repo.ITransactionRepo
}

func NewTransactionUsecase(repo repo.ITransactionRepo) ITransactionUsecase {
	return &TransactionUsecase{
		repo: repo,
	}
}

func (r TransactionUsecase) GetAll() (*dto.TransactionListResponse, error) {

	data, err := r.repo.FetchAll()
	if err != nil {
		return nil, err
	}

	response := dto.NewTransactionListResponse(data)

	return response, nil
}
