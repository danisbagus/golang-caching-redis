package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/danisbagus/golang-caching-redis/internal/handler"
	"github.com/danisbagus/golang-caching-redis/internal/repo"
	"github.com/danisbagus/golang-caching-redis/internal/usecase"
	_redis "github.com/danisbagus/golang-caching-redis/pkg/redis"

	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mysqlClient := GetMysqlClient()

	router := mux.NewRouter()

	redisConn := GetRedisConnection()
	redisClient := _redis.NewRedisClient(redisConn)

	transactionRepo := repo.NewTransactionRepo(mysqlClient)
	transactionUsecase := usecase.NewTransactionUsecase(transactionRepo, redisClient)
	transactionHandler := handler.NewTransactionHandler(transactionUsecase)

	router.HandleFunc("/api/transactions", transactionHandler.GetAllTransaction).Methods(http.MethodGet)

	appPort := "localhost:7000"
	fmt.Println("Starting the application at:", appPort)
	log.Fatal(http.ListenAndServe(appPort, router))

}

func GetMysqlClient() *sqlx.DB {
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

func GetRedisConnection() redis.Conn {
	pool := &redis.Pool{
		MaxActive:   10,
		MaxIdle:     10,
		Wait:        true,
		IdleTimeout: 5 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "localhost:7062")
			if err != nil {
				panic(err)
			}
			if _, err := c.Do("AUTH", "mypass"); err != nil {
				c.Close()
				panic(err)
			}
			return c, err
		},
	}

	conn := pool.Get()

	if r, _ := redis.String(conn.Do("PING")); r != "PONG" {
		if err := conn.Close(); err != nil {
			panic(err)
		}

		panic(errors.New("redis connect failed"))
	}

	return conn
}
