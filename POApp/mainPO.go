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

	// START MSSQL: Connections
	connString := fmt.Sprintf("server=%s;sa port=%d;database=%s;encrypt=disable", server, port, database)

	db, err = sqlx.Open("mssql", connString)
	ctx := context.Background()
	err = db.PingContext(ctx)

	if err != nil {
		log.Fatal(" Error open db:", err.Error())
	}

	fmt.Printf("Connected!\n")

	// DATE format
	datetime := time.Now().Format("2006-01-02")	
	// datetime = "2021-04-10"
	poSQL := fmt.Sprintf("SELECT * FROM fss.dbo.bsPO WHERE EditDate >= '%s 00:00:00' ORDER BY EditDate DESC;", datetime)

	poStore, err := metaapis.ReadPOData(db, poSQL)

	if err != nil {
		log.Fatal("Error reading PO: ", err.Error())
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
	// Genesis Checking
	_, _ = metaapis.ReadCloud("SO", ctx, client, true)

	// Adding data to cloud
	metaapis.AddCloudPO(ctx, client, poStore)

	// END FIREBASE: fIRESTORE

	fmt.Println("Runtime: ", time.Since(runStart))
}

// END: MAIN
