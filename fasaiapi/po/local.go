package fasaiapi

import (
	"context"
	"database/sql"
	"fmt"
)

// PO type structure data model from local database
type PO struct {
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
	TotalAmt     string
	DiscountAmt  string
	BefoeTaxAmt  string
	NetAmt       string
	DueDate      string
	CompleteDate sql.NullString
	CreateDate   string
	CancelDate   sql.NullString
	EditDate     sql.NullString
}

// ReadPOLocal get po data from local database
func ReadPOLocal(db *sql.DB, datetime string, isGenesis bool) (map[string]PO, error) {
	fields := fmt.Sprint("DocNo, DocDate, RefNo, RefDate, PoNo, Status, TaxType, ApID, ApName, Credit, TotalAmt, DiscountAmt, BefoeTaxAmt, NetAmt, DueDate, CompleteDate, CreateDate, CancelDate, EditDate")
	statementPO := fmt.Sprintf("SELECT %[1]s FROM fss.dbo.bsPO WHERE DocDate >= '%[2]s 00:00:00' AND EditDate <= '%[2]s 20:00:00' ORDER BY EditDate DESC", fields, datetime)

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
			&poRow.DocNo,
			&poRow.DocDate,
			&poRow.RefNo,
			&poRow.RefDate,
			&poRow.PoNo,
			&poRow.Status,
			&poRow.TaxType,
			&poRow.ApID,
			&poRow.ApName,
			&poRow.Credit,
			&poRow.TotalAmt,
			&poRow.DiscountAmt,
			&poRow.BefoeTaxAmt,
			&poRow.NetAmt,
			&poRow.DueDate,
			&poRow.CompleteDate,
			&poRow.CreateDate,
			&poRow.CancelDate,
			&poRow.EditDate,
		)

		if err != nil {
			return store, err
		}

		store[poRow.DocNo] = poRow

		// IF YOU WANT TO SEE DATA STREAM FROM LOCAL DATABASE
		fmt.Printf("DocNo: %s, PoNo: %s, NetAmt: %s\n", poRow.DocNo, poRow.PoNo, poRow.NetAmt)
	}

	return map[string]PO{}, nil
}
