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
	fmt.Printf("Connected!\n")

	if err != nil {
		fmt.Println(" Error open db:", err.Error())
	}

	// DATE format
	datetime := time.Now().Format("2006-01-02")
	// DEBUG: 
	datetime = "2021-04-10"
	soSQL := fmt.Sprintf("SELECT * FROM fss.dbo.bsSaleOrder WHERE EditDate >= '%s 00:00:00' ORDER BY EditDate DESC;", datetime)
	soItemSQL := fmt.Sprintf("SELECT * FROM fss.dbo.bsSaleOrderItem WHERE EditDate >= '%s 00:00:00' ORDER BY EditDate DESC;", datetime)

	soStore, errSo := metaapis.ReadSOData(db, soSQL)
	soItemStore, errSoItem := metaapis.ReadSOItemData(db, soItemSQL)

	if errSo != nil {
		log.Fatal("Error reading SO: ", errSo.Error())
	}

	if errSoItem != nil {
		log.Fatal("Error reading SO Item: ", errSoItem.Error())
	}

	defer db.Close()

	log.Println(soItemStore)

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

	// ADDING OR INIT DATA
	_, _ = metaapis.ReadCloud("SO", ctx, client, true)
	_, _ = metaapis.ReadCloud("SOItem", ctx, client, true)

	// Add SO to cloud
	metaapis.AddCloudSO(ctx, client, soStore)

	fmt.Println("Runtime: ", time.Since(runStart))
}

// END: MAIN
