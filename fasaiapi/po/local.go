package fasaiapi

import (
	"context"
	"database/sql"
	"fmt"
)

// PO type structure data model from local database
type PO struct {
	name         string
	DocNo        string
	DocDate      string
	RefNo        string
	RefDate      string
	PoNo         string
	Status       string
	TaxType      string
	ApID         string
	ApName       string
	Credit       string
	DueDate      string
	CompleteDate string
	PayType      string
	TotalAmt     string
	DiscountText string
	DiscountAmt  string
	BefoeTaxAmt  string
	NetAmt       string
	RemainAmt    string
	Description  string
	CreateBy     string
	CreateDate   string
	CancelBy     string
	CancelDate   string
	EditDate     string
}

// ReadPOLocal get po data from local database
func ReadPOLocal(db *sql.DB, datetime string, isGenesis bool) (map[string]PO, error) {
	statementPO := fmt.Sprintf("SELECT * FROM fss.dbo.bsPR WHERE DocDate >= '%[1]s 00:00:00' AND DocDate <= '%[1]s 20:00:00' ORDER BY DocDate DESC", datetime)

	store := make(map[string]PO)
	ctx := context.Background()

	// CHECK IF DATABASE IS ALIVE.
	err := db.PingContext(ctx)
	if err != nil {
		return store, err
	}

		// EXECUTE QUERY
	rows, err := db.QueryContext(ctx, statementPO)
	if err != nil {
		return store, err
	}

	// CLOSE CONNECTION
	defer rows.Close()

	// ITERATE THROUGH THE RESULT SET.
	for rows.Next() {
		var poRow PO

		err := rows.Scan(
			
		)
	}

	return map[string]PO{}, nil
}
