package metaapis

import (
	"context"
	"database/sql"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	models "DataFlow/models"
)

var err error

// AddStocks Add new stock data to cloud in case that database doesn't not exist.
func AddCloudStocks(ctx context.Context, client *firestore.Client, storeData map[string]models.Stock) {
	if len(storeData) == 0 {
		log.Println("Stock: Up-to-date.")
		return
	}

	var err error
	for key, data := range storeData {
		_, err = client.Collection("Stocks").Doc(key).Set(ctx, map[string]interface{}{
			"ID":           data.ID,
			"Name":         data.Name,
			"GroupID":      data.GroupID,
			"Cost":         []string{data.Cost},
			"Price":        []string{data.Price},
			"StockQty":     []string{data.StockQty},
			"StockValue":   []string{data.StockValue},
			"LastBuyDate":  []sql.NullString{data.LastBuyDate},
			"LastSellDate": []sql.NullString{data.LastSellDate},
			"EditDate":     []sql.NullString{data.EditDate},
		})
	}

	if err != nil {
		log.Fatalf("Failed adding Stock type: %v", err)
	}

	log.Println("Completed Adding Stock to cloud.")
}

// ReadStock get data from cloud
func ReadCloudStock(ctx context.Context, client *firestore.Client) []map[string]interface{} {
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

// UpdateStock to modify cloud database
func UpdateCloudStock(ctx context.Context, client *firestore.Client, stockID string, store map[string]interface{}) {
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

// PrepareAndUpdateStocks to adjust data format and upload to cloud
func PrepareAndUpdateStocks(ctx context.Context, client *firestore.Client, cloudDB []map[string]interface{}, localDB map[string]models.Stock) map[string]interface{} {
	updatedStore := make(map[string]interface{})

	for _, cdata := range cloudDB {
		stockID := cdata["ID"].(string)
		stockPrices := cdata["Price"].([]interface{})
		stockCosts := cdata["Cost"].([]interface{})
		stockQties := cdata["StockQty"].([]interface{})
		stockValues := cdata["StockValue"].([]interface{})
		stockLastBuyDates := cdata["LastBuyDate"].([]interface{})
		stockLastSellDates := cdata["LastSellDate"].([]interface{})
		stockEditDates := cdata["EditDate"].([]interface{})

		localData := localDB[stockID]
		if localData.Price == "" {
			continue
		}

		stockPrices = append(stockPrices, localData.Price)
		stockCosts = append(stockCosts, localData.Cost)
		stockQties = append(stockQties, localData.StockQty)
		stockValues = append(stockValues, localData.StockValue)
		stockLastBuyDates = append(stockLastBuyDates, localData.LastBuyDate)
		stockLastSellDates = append(stockLastSellDates, localData.LastSellDate)
		stockEditDates = append(stockEditDates, localData.EditDate)

		updatedStore["stockPrices"] = stockPrices
		updatedStore["stockCosts"] = stockCosts
		updatedStore["stockQties"] = stockQties
		updatedStore["stockValues"] = stockValues
		updatedStore["stockLastBuyDates"] = stockLastBuyDates
		updatedStore["stockLastSellDates"] = stockLastSellDates
		updatedStore["stockEditDates"] = stockEditDates

		UpdateCloudStock(ctx, client, stockID, updatedStore)
	}

	return updatedStore
}
