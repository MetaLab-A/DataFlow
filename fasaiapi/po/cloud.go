package fasaiapi

import (
	"context"
	"database/sql"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

var err error

// AddPO Add new PO data to cloud in case that database doesn't not exist.
func AddPO(ctx context.Context, client *firestore.Client, poData map[string]PO) {
	for key, data := range poData {
		_, err = client.Collection("Stocks").Doc(key).Set(ctx, map[string]interface{}{
			"DocNo":      data.DocNo,
			"DocDate":      data.DocDate,
			"RefNo":        data.RefNo,
			"RefDate":      data.RefDate,
			"PoNo":         data.PoNo,
			"Status":       data.Status,
			"TaxType":      data.TaxType,
			"ApID":         data.ApName,
			"ApName":       data.ApName,
			"Credit":       data.Credit,
			"TotalAmt":     data.TotalAmt,
			"DiscountAmt":  data.DiscountAmt,
			"BefoeTaxAmt":  data.BefoeTaxAmt,
			"NetAmt":       data.NetAmt,
			"DueDate":      data.DueDate,
			"CompleteDate": data.CompleteDate,
			"CreateDate":   data.CreateDate,
			"CancelDate":   data.CancelDate,
			"EditDate":     data.EditDate,
		})
	}

	if err != nil {
		log.Fatalf("Failed adding Stock type: %v", err)
	}
}