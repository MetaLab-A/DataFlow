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

	"DataFlow/localsql"
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
	datetime = "2021-02-06"
	stockStore, err := localsql.ReadStockSQL(db, datetime)
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
	// cloudDB := readStock(ctx)

	// for _, data := range cloudDB {
	// 	stockID := data["id"].(string)
	// 	// localData := stockStore[stockID]
	// 	// stockPrices := data["price"]
	// 	// stockCosts := data["cost"]

	// 	// updateStock(ctx, stockID, "123", "321")
	// }

	fmt.Println("Runtime: ", time.Since(runStart))
}

func addStocks(ctx context.Context, stockData map[string]localsql.Stock) {
	for key, data := range stockData {
		_, err = client.Collection("Stocks").Doc(key).Set(ctx, data)
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

func updateStock(ctx context.Context, stockID string, newPrice []string, newCost []string) {
	_, err = client.Collection("Stocks").Doc(stockID).Update(ctx, []firestore.Update{
		{
			Path:  "price",
			Value: newPrice,
		},
		{
			Path:  "cost",
			Value: newCost,
		},
	})

	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	} else {
		log.Printf("Stocks ID: %s Updated", stockID)
	}
}
