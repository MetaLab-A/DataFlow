package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	_ "github.com/denisenkom/go-mssqldb"
	"google.golang.org/api/option"

	fsStocks "DataFlow/fasaiapi/stocks"
)

var db *sql.DB
var server = "(local)"
var port = 1433
var database = "fss"
var err error
var client *firestore.Client

// START: MAIN
func main() {
	runStart := time.Now()
	
	// curTime = runStart.Format("15:04")

	// START MSSQL: Connections
	connString := fmt.Sprintf("server=%s;sa port=%d;database=%s;encrypt=disable", server, port, database)

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

	// DATE format
	// datetime = "2021-02-06"
	datetime := time.Now().Format("2006-01-02")

	stockStore, err := fsStocks.ReadStockLocal(db, datetime, false)

	if err != nil {
		log.Fatal("Error reading Stock: ", err.Error())
	}

	defer db.Close()

	// END MSSQL: Connections

	// START FIREBASE: fIRESTORE
	sa := option.WithCredentialsFile("fasai-cloud-firebase-adminsdk-iu86z-5d3ce4573f.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err = app.Firestore(ctx)
	defer client.Close()

	if err != nil {
		log.Fatalln(err)
	}
	// END FIREBASE: fIRESTORE

	// START: Stocks part
	// ADDING OR INIT DATA
	// addStocks(ctx, stockStore)

	cloudDB := fsStocks.ReadStock(ctx, client)
	fsStocks.PrepareAndUpdateStocks(ctx, client, cloudDB, stockStore)
	// END: Stocks part

	fmt.Println("Runtime: ", time.Since(runStart))
}

// END: MAIN
