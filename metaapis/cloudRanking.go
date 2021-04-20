package metaapis

import (
	models "DataFlow/models"
	"context"
	"log"

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
	}

	if err != nil {
		log.Fatalf("Failed adding Ranking Item type: %v", err)
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
			"ItemID":     data.ItemID,
			"ItemName":   data.ItemName,
			"HighPrice":  data.HighPrice,
			"LowPrice":   data.LowPrice,
			"Qty":        data.Qty,
		})

		log.Println("Ranking SOItem:", data.ItemID, "Added")
	}

	if err != nil {
		log.Fatalf("Failed adding Ranking SOItem type: %v", err)
	}

	log.Println("Completed Adding Ranking SOItem to cloud.")
}
