package main

import (
	"context"
	"fmt"
	"log"
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

	// curTime = runStart.Format("15:04")

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

	// DATE format
	datetime := time.Now().Format("2006-01-02")
	// DEBUG:
	datetime = "2021-04-01"
	invItemSQL := fmt.Sprintf("SELECT * FROM fss.dbo.bsInvoiceItem WHERE EditDate >= '%s 00:00:00' AND DocNo LIKE 'VS%%' ORDER BY EditDate DESC;", datetime)

	invItemStore, errSoItem := metaapis.ReadInvoiceItemData(db, invItemSQL)

	if errSoItem != nil {
		log.Fatal("Error reading SO Item: ", errSoItem.Error())
	}

	defer db.Close()

	rankingStore := calInvItem2RankingItem(invItemStore)
	printRanking(rankingStore)

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
	// END FIREBASE: fIRESTORE

	fmt.Println("Runtime: ", time.Since(runStart))
}

// END: MAIN

func calInvItem2RankingItem(store map[string]models.InvoiceItem) map[string]*models.RankingItem {
	repeatedID := make(map[string]*models.RankingItem)

	for _, s := range store {
		tempObj := repeatedID[s.ItemID]
		fTotal, _ := strconv.ParseFloat(s.TotalAmt, 64)
		fProfit, _ := strconv.ParseFloat(s.ProfitAmt, 64)
		fQty, _ := strconv.Atoi(s.Qty)
		fPrice, _ := strconv.ParseFloat(s.Price, 64)

		// Create new data in map if it found first time
		if tempObj == nil {
			tempObj = &models.RankingItem{}
		}

		// High - Low decision making
		tempObj.HighPrice = newHigh(fPrice, tempObj.HighPrice)
		tempObj.LowPrice = newLow(fPrice, tempObj.HighPrice)

		tempObj.TotalAmt += fTotal
		tempObj.ProfitAmt += fProfit
		tempObj.Qty += fQty
		repeatedID[s.ItemID] = tempObj
	}

	return repeatedID
}

func newHigh(incoming float64, record float64) float64 {
	if record <= 0 || incoming > record {
		return incoming
	}
	return record
}

func newLow(incoming float64, record float64) float64 {
	if record <= 0 || incoming < record {
		return incoming
	}
	return record
}

func printRanking(store map[string]*models.RankingItem) {
	for k, v := range store {
		fmt.Println(k, ":", v)
	}
}
