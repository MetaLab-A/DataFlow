package metaapis

import (
	"context"
	"log"	
	models "DataFlow/models"
	"cloud.google.com/go/firestore"
)

// AddPO Add new PO data to cloud in case that database doesn't not exist.
func AddCloudPO(ctx context.Context, client *firestore.Client, poData map[string]models.PO) {
	var err error

	for key, data := range poData {
		_, err = client.Collection("PO").Doc(key).Set(ctx, map[string]interface{}{
			"RowOrder": data.RowOrder,
			"DocNo": data.DocNo,
			"DocDate": data.DocDate,
			"DocType": data.DocType,
			"RefNo": data.RefNo,
			"RefDate": data.RefDate,
			"PoNo": data.PoNo,
			"Status": data.Status,
			"TaxType": data.TaxType,
			"ApID": data.ApID,
			"ApName": data.ApName,
			"Credit": data.Credit,
			"DueDate": data.DueDate,
			"CompleteDate": data.CompleteDate,
			"PayType": data.PayType,
			"TotalAmt": data.TotalAmt,
			"DiscountText": data.DiscountText,
			"DiscountAmt": data.DiscountAmt,
			"BefoeTaxAmt": data.BefoeTaxAmt,
			"TaxAmt": data.TaxAmt,
			"TaxRate": data.TaxRate,
			"NetAmt": data.NetAmt,
			"RemainAmt": data.RemainAmt,
			"Description": data.Description,
			"CreateBy": data.CreateBy,
			"CreateDate": data.CreateDate,
			"CancelBy": data.CancelBy,
			"CancelDate": data.CancelDate,
			"AddID": data.AddID,
			"AddDate": data.AddDate,
			"EditID": data.EditID,
			"EditDate": data.EditDate,
		})
		log.Println("Added PO:", data.DocNo)
	}

	if err != nil {
		log.Fatalf("Failed adding PO type: %v", err)
	}

	log.Println("Completed Adding PO to cloud.")
}
