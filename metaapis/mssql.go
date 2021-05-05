package metaapis

import (
	"context"
	_ "fmt"
	models "DataFlow/models"
	sqlx "github.com/jmoiron/sqlx"
)

func ReadStockData(db *sqlx.DB, statementSQL string) (map[string]models.Stock, error) {
	store := make(map[string]models.Stock)
	ctx := context.Background()

	// CHECK IF DATABASE IS ALIVE.
	err := db.PingContext(ctx)
	if err != nil {
		return store, err
	}
	// EXECUTE QUERY
	rows, err := db.Queryx(statementSQL)
	if err != nil {
		return store, err
	}
	// CLOSE CONNECTION
	defer rows.Close()
	// ITERATE THROUGH THE RESULT SET.
	for rows.Next() {
		// GET VALUES FROM ROW.
		var model models.Stock
		err := rows.StructScan(&model)

		if err != nil {
			return store, err
		}
		store[model.ID] = model
	}
	return store, nil
}

func ReadSOData(db *sqlx.DB, statementSQL string) (map[string]models.SO, error) {
	store := make(map[string]models.SO)
	ctx := context.Background()

	// CHECK IF DATABASE IS ALIVE.
	err := db.PingContext(ctx)
	if err != nil {
		return store, err
	}
	// EXECUTE QUERY
	rows, err := db.Queryx(statementSQL)
	if err != nil {
		return store, err
	}
	// CLOSE CONNECTION
	defer rows.Close()
	// ITERATE THROUGH THE RESULT SET.
	for rows.Next() {
		// GET VALUES FROM ROW.
		var model models.SO
		err := rows.StructScan(&model)

		if err != nil {
			return store, err
		}
		store[model.DocNo] = model
	}
	return store, nil
}

func ReadSOItemData(db *sqlx.DB, statementSQL string) (map[string]models.SOItem, error) {
	store := make(map[string]models.SOItem)
	ctx := context.Background()

	// CHECK IF DATABASE IS ALIVE.
	err := db.PingContext(ctx)
	if err != nil {
		return store, err
	}
	// EXECUTE QUERY
	rows, err := db.Queryx(statementSQL)
	if err != nil {
		return store, err
	}
	// CLOSE CONNECTION
	defer rows.Close()
	// ITERATE THROUGH THE RESULT SET.
	for rows.Next() {
		// GET VALUES FROM ROW.
		var model models.SOItem
		err := rows.StructScan(&model)

		if err != nil {
			return store, err
		}
		store[model.RowOrder] = model
	}
	return store, nil
}

func ReadPOData(db *sqlx.DB, statementSQL string) (map[string]models.PO, error) {
	store := make(map[string]models.PO)
	ctx := context.Background()

	// CHECK IF DATABASE IS ALIVE.
	err := db.PingContext(ctx)
	if err != nil {
		return store, err
	}
	// EXECUTE QUERY
	rows, err := db.Queryx(statementSQL)
	if err != nil {
		return store, err
	}
	// CLOSE CONNECTION
	defer rows.Close()
	// ITERATE THROUGH THE RESULT SET.
	for rows.Next() {
		// GET VALUES FROM ROW.
		var model models.PO
		err := rows.StructScan(&model)

		if err != nil {
			return store, err
		}
		store[model.DocNo] = model
	}
	return store, nil
}

func ReadPOItemData(db *sqlx.DB, statementSQL string) (map[string]models.POItem, error) {
	store := make(map[string]models.POItem)
	ctx := context.Background()

	// CHECK IF DATABASE IS ALIVE.
	err := db.PingContext(ctx)
	if err != nil {
		return store, err
	}
	// EXECUTE QUERY
	rows, err := db.Queryx(statementSQL)
	if err != nil {
		return store, err
	}
	// CLOSE CONNECTION
	defer rows.Close()
	// ITERATE THROUGH THE RESULT SET.
	for rows.Next() {
		// GET VALUES FROM ROW.
		var model models.POItem
		err := rows.StructScan(&model)

		if err != nil {
			return store, err
		}
		store[model.RowOrder] = model
	}
	return store, nil
}


func ReadInvoiceItemData(db *sqlx.DB, statementSQL string) (map[string]models.InvoiceItem, error) {
	store := make(map[string]models.InvoiceItem)
	ctx := context.Background()

	// CHECK IF DATABASE IS ALIVE.
	err := db.PingContext(ctx)
	if err != nil {
		return store, err
	}
	// EXECUTE QUERY
	rows, err := db.Queryx(statementSQL)
	if err != nil {
		return store, err
	}
	// CLOSE CONNECTION
	defer rows.Close()
	// ITERATE THROUGH THE RESULT SET.
	for rows.Next() {
		// GET VALUES FROM ROW.
		var model models.InvoiceItem
		err := rows.StructScan(&model)

		if err != nil {
			return store, err
		}
		store[model.RowOrder] = model
	}
	return store, nil
}


func ReadVSSummary(db *sqlx.DB) (map[string]models.VSSummary, error) {
	store := make(map[string]models.VSSummary)
	ctx := context.Background()
	statementSQL := "SELECT ItemID, SUM(Qty) AS Qty, SUM(TotalAmt) AS TotalAmt FROM fss.dbo.bsInvoiceItem WHERE AddDate >= '2021-01-01' AND DocNo LIKE 'VS%%' GROUP BY ItemID"

	// CHECK IF DATABASE IS ALIVE.
	err := db.PingContext(ctx)
	if err != nil {
		return store, err
	}
	// EXECUTE QUERY
	rows, err0 := db.Queryx(statementSQL)

	if err != nil {
		return store, err0
	}
	// CLOSE CONNECTION
	defer rows.Close()

	// ITERATE THROUGH THE RESULT SET.
	for rows.Next() {
		// GET VALUES FROM ROW.
		var model models.VSSummary
		err := rows.StructScan(&model)

		if err != nil {
			return store, err
		}
		store[model.ItemID] = model
	}
	return store, nil
}


func ReadItemName(db *sqlx.DB) (map[string]models.ItemName, error) {
	store := make(map[string]models.ItemName)
	ctx := context.Background()
	statementSQL := "SELECT ItemID, ItemName FROM fss.dbo.bsInvoiceItem WHERE AddDate >= '2021-01-01' AND DocNo LIKE 'VS%%'"

	// CHECK IF DATABASE IS ALIVE.
	err := db.PingContext(ctx)
	if err != nil {
		return store, err
	}
	// EXECUTE QUERY
	rows, err0 := db.Queryx(statementSQL)

	if err != nil {
		return store, err0
	}
	// CLOSE CONNECTION
	defer rows.Close()

	// ITERATE THROUGH THE RESULT SET.
	for rows.Next() {
		// GET VALUES FROM ROW.
		var model models.ItemName
		err := rows.StructScan(&model)

		if err != nil {
			return store, err
		}
		store[model.ItemID] = model
	}
	return store, nil
}

func ReadSOSummary(db *sqlx.DB) (map[string]models.SOItem, error) {
	store := make(map[string]models.SOItem)
	ctx := context.Background()
	statementSQL := "SELECT ItemID, SUM(Qty) AS Qty, SUM(RQty) AS RQty FROM fss.dbo.bsSaleOrderItem WHERE AddDate >= '2021-01-01' AND DocNo LIKE 'SO%' GROUP BY ItemID"

	// CHECK IF DATABASE IS ALIVE.
	err := db.PingContext(ctx)
	if err != nil {
		return store, err
	}
	// EXECUTE QUERY
	rows, err0 := db.Queryx(statementSQL)

	if err != nil {
		return store, err0
	}
	// CLOSE CONNECTION
	defer rows.Close()

	// ITERATE THROUGH THE RESULT SET.
	for rows.Next() {
		// GET VALUES FROM ROW.
		var model models.SOItem
		err := rows.StructScan(&model)

		if err != nil {
			return store, err
		}
		store[model.ItemID] = model
	}
	return store, nil
}

func ReadStockSummary(db *sqlx.DB) (map[string]models.Stock, error) {
	store := make(map[string]models.Stock)
	ctx := context.Background()

	// Fix here
	statementSQL := "SELECT ID, StockQty FROM fss.dbo.bsItem WHERE EditDate >= '2021-01-01'"

	// CHECK IF DATABASE IS ALIVE.
	err := db.PingContext(ctx)
	if err != nil {
		return store, err
	}
	// EXECUTE QUERY
	rows, err := db.Queryx(statementSQL)

	if err != nil {
		return store, err
	}
	// CLOSE CONNECTION
	defer rows.Close()

	// ITERATE THROUGH THE RESULT SET.
	for rows.Next() {
		// GET VALUES FROM ROW.
		var model models.Stock
		err := rows.StructScan(&model)

		if err != nil {
			return store, err
		}
		store[model.ID] = model
	}
	return store, nil
}

func ReadPOSummary(db *sqlx.DB) (map[string]models.POItem, error) {
	store := make(map[string]models.POItem)
	ctx := context.Background()
	statementSQL := "SELECT ItemID, SUM(Qty) AS Qty, SUM(RQty) AS RQty FROM fss.dbo.bsPRItem where DocNo like 'PO%' and AddDate >= '2021-01-01' GROUP BY ItemID"

	// CHECK IF DATABASE IS ALIVE.
	err := db.PingContext(ctx)
	if err != nil {
		return store, err
	}
	// EXECUTE QUERY
	rows, err := db.Queryx(statementSQL)

	if err != nil {
		return store, err
	}
	// CLOSE CONNECTION
	defer rows.Close()

	// ITERATE THROUGH THE RESULT SET.
	for rows.Next() {
		// GET VALUES FROM ROW.
		var model models.POItem
		err := rows.StructScan(&model)

		if err != nil {
			return store, err
		}
		store[model.ItemID] = model
	}
	return store, nil
}

func ReadRRSummary(db *sqlx.DB) (map[string]models.POItem, error) {
	store := make(map[string]models.POItem)
	ctx := context.Background()
	statementSQL := "SELECT ItemID, SUM(Qty) AS Qty FROM fss.dbo.bsPOItem where DocNo like 'RR%' and AddDate >= CAST(GETDATE() as date) and AddDate < CAST(GETDATE()+1 as date) GROUP BY ItemID"

	// CHECK IF DATABASE IS ALIVE.
	err := db.PingContext(ctx)
	if err != nil {
		return store, err
	}
	// EXECUTE QUERY
	rows, err := db.Queryx(statementSQL)

	if err != nil {
		return store, err
	}
	// CLOSE CONNECTION
	defer rows.Close()

	// ITERATE THROUGH THE RESULT SET.
	for rows.Next() {
		// GET VALUES FROM ROW.
		var model models.POItem
		err := rows.StructScan(&model)

		if err != nil {
			return store, err
		}
		store[model.ItemID] = model
	}
	return store, nil
}
