package fasaiapi

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)


// AddPO Add new PO data to cloud in case that database doesn't not exist.
func AddPO(ctx context.Context, client *firestore.Client, poData map[string]PO) {
	var err error

	for key, data := range poData {
		_, err = client.Collection("PO").Doc(key).Set(ctx, map[string]interface{}{
			"DocNo":        data.DocNo,
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
		log.Fatalf("Failed adding PO type: %v", err)
	}

	log.Println("Added PO data to Firebase")
}
