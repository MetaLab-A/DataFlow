package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	_ "github.com/denisenkom/go-mssqldb"
	"google.golang.org/api/option"

	metaapis "DataFlow/metaapis"
	sqlx "github.com/jmoiron/sqlx"
)

var db *sqlx.DB
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

	db, err = sqlx.Open("mssql", connString)
	ctx := context.Background()
	err = db.PingContext(ctx)

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Database Connected")
	
	if err != nil {
		fmt.Println(" Error open db:", err.Error())
	}

	// DATE format
	datetime := time.Now().Format("2006-01-02")
	stockSQL := fmt.Sprintf("SELECT * FROM fss.dbo.bsItem WHERE EditDate >= '%s 00:00:00' AND GroupID IN ('C', 'C-1', 'E') ORDER BY EditDate DESC;", datetime)

	stockStore, err := metaapis.ReadStockData(db, stockSQL)

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

	cloudDB, _ := metaapis.ReadCloud("Stocks", ctx, client, false)
	_ = metaapis.PrepareAndUpdateStocks(ctx, client, cloudDB, stockStore)
	// END: Stocks part

	fmt.Println("Runtime: ", time.Since(runStart))
	time.Sleep(2 * time.Second)
}

// END: MAIN
