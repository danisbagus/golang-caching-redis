package handler

import (
	"encoding/json"
	"net/http"

	"github.com/danisbagus/golang-caching-redis/internal/usecase"
)

type TransactionHandler struct {
	usecase usecase.ITransactionUsecase
}

func NewTransactionHandler(usecase usecase.ITransactionUsecase) *TransactionHandler {
	return &TransactionHandler{
		usecase: usecase,
	}
}

func (rc TransactionHandler) GetAllTransaction(w http.ResponseWriter, r *http.Request) {
	dataList, err := rc.usecase.GetAll()
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, err)
		return
	}
	writeResponse(w, http.StatusOK, dataList)
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
