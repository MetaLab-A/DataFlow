package metaapis

import (
	"context"
	"log"

	models "DataFlow/models"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

// AddStocks Add new stock data to cloud in case that database doesn't not exist.
func AddCloudSO(ctx context.Context, client *firestore.Client, storeData map[string]models.SO) {
	for key, data := range storeData {
		_, err = client.Collection("SO").Doc(key).Set(ctx, map[string]interface{}{
			"RowOrder":     data.RowOrder,
			"DocNo":        data.DocNo,
			"DocDate":      data.DocDate,
			"DocType":      data.DocType,
			"RefNo":        data.RefNo,
			"Status":       data.Status,
			"TaxType":      data.TaxType,
			"IsNotify":     data.IsNotify,
			"ArID":         data.ArID,
			"ArName":       data.ArName,
			"AreaID":       data.AreaID,
			"SalemanID":    data.SalemanID,
			"TeamID":       data.TeamID,
			"Credit":       data.Credit,
			"DueDate":      data.DueDate,
			"DeliveryDate": data.DeliveryDate,
			"CompleteDate": data.CompleteDate,
			"PayType":      data.PayType,
			"TotalAmt":     data.TotalAmt,
			"DiscountText": data.DiscountText,
			"DiscountAmt":  data.DiscountAmt,
			"BefoeTaxAmt":  data.BefoeTaxAmt,
			"TaxAmt":       data.TaxAmt,
			"TaxRate":      data.TaxRate,
			"NetAmt":       data.NetAmt,
			"RemainAmt":    data.RemainAmt,
			"Description":  data.Description,
			"CreateBy":     data.CreateBy,
			"CreateDate":   data.CreateDate,
			"CancelBy":     data.CancelBy,
			"CancelDate":   data.CancelDate,
			"AddID":        data.AddID,
			"AddDate":      data.AddDate,
			"EditID":       data.EditDate,
			"EditDate":     data.EditDate,
		})
	}

	if err != nil {
		log.Fatalf("Failed adding SO type: %v", err)
	}
}

// ReadStock get data from cloud
func ReadCloudSO(ctx context.Context, client *firestore.Client) []map[string]interface{} {
	store := make([]map[string]interface{}, 0)
	iter := client.Collection("SO").Documents(ctx)

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
func UpdateCloudSO(ctx context.Context, client *firestore.Client, stockID string, store map[string]interface{}) {
	_, err = client.Collection("SO").Doc(stockID).Update(ctx, []firestore.Update{
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
func PrepareAndUpdateSO(ctx context.Context, client *firestore.Client, cloudDB []map[string]interface{}, localDB map[string]models.Stock) map[string]interface{} {
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
