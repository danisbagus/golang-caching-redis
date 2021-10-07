package dto

import (
	"github.com/danisbagus/golang-caching-redis/internal/domain"
)

type TransactionResponse struct {
	TransactionID string `json:"transaction_id"`
	MerchantID    int64  `json:"merchant_id"`
	SKUID         string `json:"sku_id"`
	SuppierID     int64  `json:"supplier_id"`
	Quantity      int64  `json:"quantity"`
	TotalPrice    int64  `json:"total_price"`
	CreatedAt     string `json:"created_at"`
}

type TransactionListResponse struct {
	Transactions []TransactionResponse
}

func NewTransactionListResponse(data []domain.Transaction) *TransactionListResponse {
	dataList := make([]TransactionResponse, len(data))

	for k, v := range data {
		dataList[k] = TransactionResponse{
			TransactionID: v.TransactionID,
			MerchantID:    v.MerchantID,
			SKUID:         v.SKUID,
			SuppierID:     v.SuppierID,
			Quantity:      v.Quantity,
			TotalPrice:    v.TotalPrice,
			CreatedAt:     v.CreatedAt,
		}
	}
	return &TransactionListResponse{Transactions: dataList}
}
