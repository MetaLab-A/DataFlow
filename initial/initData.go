package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
	"DataFlow/localsql"
)

var server = "(local)"
var port = 1433
var database = "fss"
var err error

func main() {
	ctx := context.Background()
	
	connString := fmt.Sprintf("server=%s;user port=%d;database=%s;encrypt=disable", server, port, database)

	db, err := sql.Open("mssql", connString)
	err = db.PingContext(ctx)
	defer db.Close()

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n")

	if err != nil {
		fmt.Println(" Error open db:", err.Error())
	}


	InitStockData(db)
}

// InitStockData init stock data
func InitStockData(db *sql.DB) {
	datetime := time.Now().Format("2006-01-02")
	// datetime = "2021-02-06"
	stockStore, err := localsql.ReadStockSQL(db, datetime)
	if err != nil {
		log.Fatal("Error reading Stock: ", err.Error())
	}

	fmt.Println(stockStore)
}
