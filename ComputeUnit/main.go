package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	_ "github.com/denisenkom/go-mssqldb"
	"google.golang.org/api/option"

	metaapis "DataFlow/metaapis"

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
	curDatetime := time.Now().Format("2006-01-02")
	processor("2021-01-01 00:00:00", curDatetime, "RankingAnnual")
	processor("2021-04-01 00:00:00", "2021-04-01 18:00:00", "rankingTest-1")
	processor("2021-04-02 00:00:00", "2021-04-04 00:00:00", "rankingTest-2")
}
// END: MAIN

func processor(from string, to string, collectionName string) {
	runStart := time.Now()

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

	invItemSQL := fmt.Sprintf("SELECT * FROM fss.dbo.bsInvoiceItem WHERE EditDate >= '%s' AND EditDate <= '%s' AND DocNo LIKE 'VS%%' ORDER BY EditDate DESC;", from, to)

	invItemStore, errSoItem := metaapis.ReadInvoiceItemData(db, invItemSQL)

	if errSoItem != nil {
		log.Fatal("Error reading Ranking Item: ", errSoItem.Error())
	}

	defer db.Close()

	rankingStore := metaapis.CalInvItem2RankingItem(invItemStore)
	metaapis.CalPrintRanking(rankingStore)

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

	metaapis.AddCloudRankingItem(ctx, client, rankingStore, collectionName)
	// END FIREBASE: fIRESTORE

	fmt.Println("Runtime: ", time.Since(runStart))
}

