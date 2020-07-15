package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ekozlova94/parking/internal/handler"
	"github.com/ekozlova94/parking/internal/middleware"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "subscription.db")
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	//noinspection GoUnhandledErrorResult
	defer db.Close()

	mux := http.NewServeMux()
	mux.Handle("/v1/subscription", http.HandlerFunc(handler.Subscription))
	mux.Handle("/v1/check", http.HandlerFunc(handler.Check))
	if err := http.ListenAndServe("localhost:8000", middleware.Db(db, mux)); err != nil {
		log.Fatalf("can not start: %s", err.Error())
	}
}
