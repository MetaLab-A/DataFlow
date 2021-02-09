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
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var db *sql.DB
var server = "(local)"
var port = 1433
var database = "fss"
var err error
var client *firestore.Client

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

	datetime := time.Now().Format("2006-01-02")
	// datetime = "2021-02-06"
	stockStore, err := ReadStockSQL(db, datetime)
	if err != nil {
		log.Fatal("Error reading Stock: ", err.Error())
	}

	defer db.Close()

	// Firebase: Firestore
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

	addStocks(ctx, stockStore)
	cloudDB := readStock(ctx)

	for _, data := range cloudDB {
		fmt.Println("CDB: ", data)
	}


	fmt.Println("Runtime: ", time.Since(runStart))
}

func addStocks(ctx context.Context, stockData []Stock) {
	for i := 0; i < len(stockData); i++ {
		_, err = client.Collection("Stocks").Doc(stockData[i].id).Set(ctx, map[string]interface{}{
			"id":       stockData[i].id,
			"name":     stockData[i].name,
			"groupID":  stockData[i].groupID,
			"cost":     stockData[i].cost,
			"price":    stockData[i].price,
			"editDate": stockData[i].editDate,
		})
	}

	if err != nil {
		log.Fatalf("Failed adding Stock type: %v", err)
	}
}

func readStock(ctx context.Context) []map[string]interface{} {
	store := make([]map[string]interface{}, 0)
	iter := client.Collection("Stocks").Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed reading Stock type: %v", err)
		}

		store = append(store, doc.Data())
	}

	return store
}
