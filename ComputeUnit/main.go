package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	_ "github.com/denisenkom/go-mssqldb"
	"google.golang.org/api/option"

	metaapis "DataFlow/metaapis"
	models "DataFlow/models"

	sqlx "github.com/jmoiron/sqlx"
)

var db *sqlx.DB
var server = "(local)"
var port = 1433
var database = "fss"
var err error
var client *firestore.Client

// START: MAIN
func main() {
	runStart := time.Now()
	curTime := time.Now().Format("15:04:05")

	if curTime < "08:00:00" || curTime > "19:00:00" {
		os.Exit(0)
	}

	// START MSSQL: Connections
	connString := fmt.Sprintf("server=%s;sa port=%d;database=%s;encrypt=disable", server, port, database)

	db, err = sqlx.Open("mssql", connString)
	ctx := context.Background()
	err = db.PingContext(ctx)

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Database Connected")

	if err != nil {
		fmt.Println(" Error open db:", err.Error())
	}

	itemNameStore, _ := metaapis.ReadItemName(db)
	vsStore, _ := metaapis.ReadVSSummary(db)
	soStore, _ := metaapis.ReadSOSummary(db)
	stockStore, _ := metaapis.ReadStockSummary(db)
	poStore, _ := metaapis.ReadPOSummary(db)
	rrStore, _ := metaapis.ReadRRSummary(db)
	mergeStore := make(map[string]*models.QtySummary)

	for k, v := range vsStore {
		vsQty, _ := strconv.Atoi(v.Qty)
		soQty, _ := strconv.Atoi(soStore[k].Qty)
		poQty, _ := strconv.Atoi(poStore[k].Qty)
		rrQty, _ := strconv.Atoi(rrStore[k].Qty)
		stockQty, _ := strconv.Atoi(stockStore[k].StockQty)
		totalAmt, _ := strconv.ParseFloat(v.TotalAmt, 64)
		VSSOQty, _ := strconv.Atoi(soStore[k].RQty)
		RRPOQty, _ := strconv.Atoi(poStore[k].RQty)

		if curTime > "09:00:00" && curTime < "18:00:00" {
			stockQty += rrQty - vsQty
		}

		mergeStore[k] = &models.QtySummary{
			ItemID:   v.ItemID,
			ItemName: itemNameStore[v.ItemID].ItemName,
			VSSOQty:  VSSOQty,
			RRPOQty:  RRPOQty,
			VSQty:    vsQty,
			SOQty:    soQty,
			RRQty:    rrQty,
			POQty:    poQty,
			StockQty: stockQty,
			TotalAmt: totalAmt,
		}
	}

	defer db.Close()
	// END MSSQL: Connections

	// START FIREBASE: fIRESTORE
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

	metaapis.AddCloudRankingQty(ctx, client, mergeStore, "RankingAnnual-Qty")
	metaapis.AddCloudRankingTimeStamp(ctx, client, "RankingUpdateTimeStamp")
	// END FIREBASE: fIRESTORE

	fmt.Println("Runtime: ", time.Since(runStart))

	time.Sleep(3 * time.Second)
}

// END MAIN
