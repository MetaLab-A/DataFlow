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

// MSSQL Connections
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
	stockStore, err := localsql.ReadStockSQL(db, datetime, false)
	if err != nil {
		log.Fatal("Error reading Stock: ", err.Error())
	}

	defer db.Close()

// FIREBASE: fIRESTORE
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

	// ADDING OR INIT DATA
	// addStocks(ctx, stockStore)

	cloudDB := readStock(ctx)
	updateStore := make(map[string]interface{})

	for _, cdata := range cloudDB {
		stockID := cdata["ID"].(string)
		stockPrices := cdata["Price"].([]interface{})
		stockCosts := cdata["Cost"].([]interface{})
		stockQties := cdata["StockQty"].([]interface{})
		stockValues := cdata["StockValue"].([]interface{})
		stockLastBuyDates := cdata["LastBuyDate"].([]interface{})
		stockLastSellDates := cdata["LastSellDate"].([]interface{})
		stockEditDates := cdata["EditDate"].([]interface{})

		localData := stockStore[stockID]
		if localData.Price == "" {
			lastIdx := len(stockPrices) - 1
			localData.Price = stockPrices[lastIdx].(string)
			localData.Cost = stockCosts[lastIdx].(string)
			localData.StockQty = stockQties[lastIdx].(string)
			localData.StockValue = stockValues[lastIdx].(string)
			localData.EditDate = stockEditDates[lastIdx].(string)
		}

		stockPrices = append(stockPrices, localData.Price)
		stockCosts = append(stockCosts, localData.Cost)
		stockQties = append(stockQties, localData.StockQty)
		stockValues = append(stockValues, localData.StockValue)
		stockLastBuyDates = append(stockLastBuyDates, localData.LastBuyDate)
		stockLastSellDates = append(stockLastSellDates, localData.LastSellDate)
		stockEditDates = append(stockEditDates, localData.EditDate)

		updateStore["stockPrices"] = stockPrices
		updateStore["stockCosts"] = stockCosts
		updateStore["stockQties"] = stockQties
		updateStore["stockValues"] = stockValues
		updateStore["stockLastBuyDates"] = stockLastBuyDates
		updateStore["stockLastSellDates"] = stockLastSellDates
		updateStore["stockEditDates"] = stockEditDates


		updateStock(ctx, stockID, updateStore)
	}

	fmt.Println("Runtime: ", time.Since(runStart))
}

func addStocks(ctx context.Context, stockData map[string]localsql.Stock) {
	for key, data := range stockData {
		_, err = client.Collection("Stocks").Doc(key).Set(ctx, map[string]interface{}{
			"ID":       data.ID,
			"Name":     data.Name,
			"GroupID":  data.GroupID,
			"Cost":     []string{data.Cost},
			"Price":    []string{data.Price},
			"StockQty": []string{data.StockQty},
			"StockValue": []string{data.StockValue},
			"LastBuyDate": []sql.NullString{data.LastBuyDate},
			"LastSellDate": []sql.NullString{data.LastSellDate},
			"EditDate": []string{data.EditDate},
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
		fmt.Println(doc.Data())
	}

	return store
}

func updateStock(ctx context.Context, stockID string, store map[string]interface{}) {
	_, err = client.Collection("Stocks").Doc(stockID).Update(ctx, []firestore.Update{
		{
			Path:  "Price",
			Value: store["stockPrices"],
		},
		{
			Path:  "Cost",
			Value: store["stockCosts"],
		},		
		{
			Path:  "StockQty",
			Value: store["stockQties"],
		},		
		{
			Path:  "StockValue",
			Value: store["stockValues"],
		},		
		{
			Path:  "LastBuyDate",
			Value: store["stockLastBuyDates"],
		},
		{
			Path:  "LastSellDate",
			Value: store["stockLastSellDates"],
		},
		{
			Path:  "EditDate",
			Value: store["stockEditDates"],
		},
	})

	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	} else {
		log.Printf("Stocks ID: %s Updated", stockID)
	}
}
