package fasaiapi

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"testing"
	"time"
)

var db *sql.DB
var server = "(local)"
var port = 1433
var database = "fss"
var err error

func TestReadPOLocal(t *testing.T) {
	connString := fmt.Sprintf("server=%s;sa port=%d;database=%s;encrypt=disable", server, port, database)

	db, err = sql.Open("mssql", connString)
	fmt.Println(db)

	ctx := context.Background()
	err = db.PingContext(ctx)
	defer db.Close()

	fmt.Printf("Connected DB from test\n")

	if err != nil {
		fmt.Println("Error open db:", err.Error())
	}

	datetime := time.Now().Format("2006-01-02")

	got, err := ReadPOLocal(db, datetime, false)
	want := map[string]PO{}

	if !reflect.DeepEqual(got, want) && err != nil {
		t.Errorf("got %v want %v", got, want)
	}
}
