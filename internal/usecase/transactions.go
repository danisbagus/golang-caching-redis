package usecase

import (
	"encoding/json"

	"github.com/danisbagus/golang-caching-redis/internal/dto"
	"github.com/danisbagus/golang-caching-redis/internal/repo"
	"github.com/danisbagus/golang-caching-redis/pkg/redis"
)

type ITransactionUsecase interface {
	GetAll() (*dto.TransactionListResponse, error)
}

type TransactionUsecase struct {
	repo  repo.ITransactionRepo
	redis redis.IRedisClient
}

func NewTransactionUsecase(repo repo.ITransactionRepo, redis redis.IRedisClient) ITransactionUsecase {
	return &TransactionUsecase{
		repo:  repo,
		redis: redis,
	}
}

func (r TransactionUsecase) GetAll() (*dto.TransactionListResponse, error) {

	dataByte, err := r.redis.GetDataRedis("GetTransactionList")
	if err == nil && dataByte != nil {
		transactionsList := new(dto.TransactionListResponse)
		if err = json.Unmarshal(dataByte, transactionsList); err != nil {
			return nil, err
		}
		return transactionsList, nil

	} else {
		transactionsList, err := r.repo.FetchAll()
		if err != nil {
			return nil, err
		}
		response := dto.NewTransactionListResponse(transactionsList)

		go r.redis.SetDataRedis("GetTransactionList", response, 60)

		return response, nil
	}
}
