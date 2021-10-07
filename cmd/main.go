package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/danisbagus/golang-caching-redis/internal/handler"
	"github.com/danisbagus/golang-caching-redis/internal/repo"
	"github.com/danisbagus/golang-caching-redis/internal/usecase"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mysqlClient := GetClient()
	router := mux.NewRouter()

	transactionRepo := repo.NewTransactionRepo(mysqlClient)
	transactionUsecase := usecase.NewTransactionUsecase(transactionRepo)
	transactionHandler := handler.NewTransactionHandler(transactionUsecase)

	router.HandleFunc("/api/transactions", transactionHandler.GetAllTransaction).Methods(http.MethodGet)

	// appPort := fmt.Sprintf("%v:%v", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))
	appPort := "localhost:7000"
	fmt.Println("Starting the application at:", appPort)
	log.Fatal(http.ListenAndServe(appPort, router))

}

func GetClient() *sqlx.DB {
	dbHost := "localhost"
	dbPort := "7060"
	dbUser := "root"
	dbPassword := "mypass"
	dbName := "gocaching"

	connection := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", dbUser, dbPassword, dbHost, dbPort, dbName)
	client, err := sqlx.Open("mysql", connection)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
