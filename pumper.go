package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

var db *sql.DB
var server = "(local)"
var port = 1433
var database = "fss"
var err error

func main() {
	runStart := time.Now()

	connString := fmt.Sprintf("server=%s;user port=%d;database=%s;encrypt=disable", server, port, database)

	db, err = sql.Open("mssql", connString)
	ctx := context.Background()
	err = db.PingContext(ctx)

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n")

	if err != nil {
		fmt.Println(" Error open db:", err.Error())
	}

	stockStore, err := ReadStock(db)
	if err != nil {
		log.Fatal("Error reading Stock: ", err.Error())
	}

	fmt.Println(stockStore[0])

	defer db.Close()

	fmt.Println("Runtime: ", time.Since(runStart))
}
