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

		log.Println("Added", data.DocNo)
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
		break
	}

	return store
}