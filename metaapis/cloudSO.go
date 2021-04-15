package metaapis

import (
	"context"
	"log"

	models "DataFlow/models"

	"cloud.google.com/go/firestore"
)

// AddCloudSO Add new SO data to cloud in case that cloud database doesn't not exist.
func AddCloudSO(ctx context.Context, client *firestore.Client, storeData map[string]models.SO) {
	if len(storeData) == 0 {
		log.Println("SO: Up-to-date.")
		return
	}

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

		log.Println("SO :", data.DocNo, "Added")
	}

	if err != nil {
		log.Fatalf("Failed adding SO type: %v", err)
	}

	log.Println("Completed Adding SO to cloud.")
}

// AddCloudSOItem Add new SO Item data to cloud in case that cloud database doesn't not exist.
func AddCloudSOItem(ctx context.Context, client *firestore.Client, storeData map[string]models.SOItem) {
		if len(storeData) == 0 {
		log.Println("SO Item: Up-to-date.")
		return
	}

	for key, data := range storeData {
		_, err = client.Collection("SOItem").Doc(key).Set(ctx, map[string]interface{}{
			"RowOrder":     data.RowOrder,
			"DocNo":        data.DocNo,
			"SequenNo":     data.SequenNo,
			"ItemID":       data.ItemID,
			"ItemName":     data.ItemName,
			"UnitID":       data.UnitID,
			"StockID":      data.StockID,
			"Price":        data.Price,
			"Cost":         data.Cost,
			"Qty":          data.Qty,
			"RQty":         data.RQty,
			"RID":          data.RID,
			"RDate":        data.RDate,
			"DlvDate":      data.DlvDate,
			"CQty":         data.CQty,
			"CID":          data.CID,
			"CDate":        data.CDate,
			"ApID":         data.ApID,
			"TrackDate":    data.TrackDate,
			"DiscountText": data.DiscountText,
			"DiscountAmt":  data.DiscountAmt,
			"TotalAmt":     data.TotalAmt,
			"Description":  data.Description,
			"AddID":        data.AddID,
			"AddDate":      data.AddDate,
			"EditID":       data.EditID,
			"EditDate":     data.EditDate,
		})

		log.Println("SO Item:", data.DocNo, "Added")
	}

	if err != nil {
		log.Fatalf("Failed adding SO Item type: %v", err)
	}

	log.Println("Completed Adding SO Item to cloud.")
}
