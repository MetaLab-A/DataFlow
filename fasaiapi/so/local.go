package fasaiapi

import (
	"context"
	"database/sql"
	"fmt"
)

// SO type structure data model from local database
type SO struct {
	DocNo        string
	DocDate      sql.NullString
	RefNo        string
	RefDate      sql.NullString
	soNo         string
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
	CreateDate   sql.NullString
	CancelDate   sql.NullString
	EditDate     sql.NullString
	SalemanID    sql.NullString
	TeamID       sql.NullString
	DeliveryDate sql.NullString
}

// ReadsoLocal get so data from local database
func ReadSOLocal(db *sql.DB, datetime string, isGenesis bool) (map[string]SO, error) {
	fields := fmt.Sprint("DocNo, DocDate, RefNo, RefDate, soNo, Status, TaxType, ApID, ApName, Credit, TotalAmt, DiscountAmt, BefoeTaxAmt, NetAmt, DueDate, CompleteDate, CreateDate, CancelDate, EditDate, SalemanID, TeamID, DeliveryDate")

	soSQL := ""

	startDate := "2021-01-01"
	endDate := datetime

	if isGenesis {
		soSQL = fmt.Sprintf("SELECT %[1]s FROM fss.dbo.bsso WHERE EditDate >= '%[2]s 00:00:00' AND EditDate <= '%[3]s 20:00:00' ORDER BY EditDate DESC", fields, startDate, endDate)
	} else {
		soSQL = fmt.Sprintf("SELECT %[1]s FROM fss.dbo.bsso WHERE EditDate >= '%[2]s 00:00:00' AND EditDate <= '%[2]s 20:00:00' ORDER BY EditDate DESC", fields, datetime)
	}

	store := make(map[string]SO)
	ctx := context.Background()

	// CHECK IF DATABASE IS ALIVE.
	err := db.PingContext(ctx)
	if err != nil {
		return store, err
	}

	// EXECUTE QUERY
	rows, err := db.QueryContext(ctx, soSQL)
	if err != nil {
		return store, err
	}

	// CLOSE CONNECTION
	defer rows.Close()

	// ITERATE THROUGH THE RESULT SET.
	for rows.Next() {
		var soRow SO

		err := rows.Scan(
			&soRow.DocNo,
			&soRow.DocDate,
			&soRow.RefNo,
			&soRow.RefDate,
			&soRow.soNo,
			&soRow.Status,
			&soRow.TaxType,
			&soRow.ApID,
			&soRow.ApName,
			&soRow.Credit,
			&soRow.TotalAmt,
			&soRow.DiscountAmt,
			&soRow.BefoeTaxAmt,
			&soRow.NetAmt,
			&soRow.DueDate,
			&soRow.CompleteDate,
			&soRow.CreateDate,
			&soRow.CancelDate,
			&soRow.EditDate,
		)

		if err != nil {
			return store, err
		}

		store[soRow.DocNo] = soRow

		// IF YOU WANT TO SEE DATA STREAM FROM LOCAL DATABASE
		fmt.Printf("DocNo: %s, soNo: %s, NetAmt: %s\n", soRow.DocNo, soRow.soNo, soRow.NetAmt)
	}

	return store, nil
}
