package metaapis

import (
	models "DataFlow/models"
	"context"
	"log"
	"time"

	"cloud.google.com/go/firestore"
)

func AddCloudRankingItem(ctx context.Context, client *firestore.Client, storeData map[string]*models.RankingItem, collection string) {
	if len(storeData) == 0 {
		log.Println("Ranking Item: Up-to-date.")
		return
	}

	for key, data := range storeData {
		_, err = client.Collection(collection).Doc(key).Set(ctx, map[string]interface{}{
			"ItemID":     data.ItemID,
			"ItemName":   data.ItemName,
			"HighPrice":  data.HighPrice,
			"LowPrice":   data.LowPrice,
			"HighCost":   data.HighCost,
			"LowCost":    data.LowCost,
			"HighMargin": data.HighMargin,
			"LowMargin":  data.LowMargin,
			"Qty":        data.Qty,
			"TotalAmt":   data.TotalAmt,
			"ProfitAmt":  data.ProfitAmt,
		})

		log.Println("Ranking Item:", data.ItemID, "Added")

		if err != nil {
			log.Printf("Failed adding Ranking Item type: %v\n", err)
		}
	}

	log.Println("Completed Adding Ranking Item to cloud.")
}

func AddCloudRankingSOItem(ctx context.Context, client *firestore.Client, storeData map[string]*models.RankingSOItem, collection string) {
	if len(storeData) == 0 {
		log.Println("Ranking SOItem: Up-to-date.")
		return
	}

	for key, data := range storeData {
		_, err = client.Collection(collection).Doc(key).Set(ctx, map[string]interface{}{
			"ItemID":    data.ItemID,
			"ItemName":  data.ItemName,
			"HighPrice": data.HighPrice,
			"LowPrice":  data.LowPrice,
			"Qty":       data.Qty,
		})

		log.Println("Ranking SOItem:", data.ItemID, "Added")

		if err != nil {
			log.Printf("Failed adding Ranking SO Item type: %v\n", err)
		}
	}

	log.Println("Completed Adding Ranking SOItem to cloud.")
}

func AddCloudRankingQty(ctx context.Context, client *firestore.Client, storeData map[string]*models.QtySummary, collection string) {
	if len(storeData) == 0 {
		log.Println("Stock Ranking Item: Up-to-date.")
		return
	}

	for key, data := range storeData {
		_, err = client.Collection(collection).Doc(key).Set(ctx, map[string]interface{}{
			"ItemID":   data.ItemID,
			"ItemName": data.ItemName,
			"VSSOQty":  data.VSSOQty,
			"RRPOQty":  data.RRPOQty,
			"VSQty":  data.SOQty,
			"SOQty":  data.SOQty,
			"RRQty":  data.RRQty,
			"POQty":  data.POQty,
			"StockQty": data.StockQty,
			"TotalAmt": data.TotalAmt,
		})

		log.Println("Qty Ranking Item:", data.ItemID, "Added")

		if err != nil {
			log.Printf("Failed adding Qty Stock Item type: %v\n", err)
		}
	}

	log.Println("Completed Adding Qty Ranking Item to cloud.")
}

func AddCloudRankingTimeStamp(ctx context.Context, client *firestore.Client, collection string) {
	updatedDatetime := time.Now().Format("2006-01-02 15:04:05")
	_, err = client.Collection(collection).Doc("UpdatedTime").Set(ctx, map[string]interface{}{
		"Timestamp": updatedDatetime,
	})

	log.Println("Completed Adding Update Timestamp to cloud.")
}
