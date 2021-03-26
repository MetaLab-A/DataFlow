package metaapis

import (
	"context"
	"database/sql"
	"fmt"
	models "DataFlow/models"
	sqlx "github.com/jmoiron/sqlx"
)


func ReadLocalData(db *sqlx.DB, statementSQL string, modelType string) (map[string]interface{}, error) {
	store := make(map[string]interface{})
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

	// CLOSE CONNECTIOn
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
		fmt.Println(model)

		// IF YOU WANT TO SEE DATA STREAM FROM LOCAL DATABASE
		// fmt.Printf("ID: %s, Name: %s, GroupID: %s, Cost: %s, Price: %s, EditDate: %s\n",
		// 	stockRow.ID, stockRow.Name, stockRow.GroupID, stockRow.Cost, stockRow.Price, stockRow.EditDate)
	}

	return store, nil
}

// ReadStockLocal READS ALL sTOCK FROM BSiTEM RECORDS
func ReadStockLocal(db *sql.DB, datetime string, isGenesis bool) (map[string]models.Stock, error) {
	store := make(map[string]models.Stock)

	ctx := context.Background()

	// CHECK IF DATABASE IS ALIVE.
	err := db.PingContext(ctx)
	if err != nil {
		return store, err
	}

	// GET CURRENT DATE
	// USE GENESIS ONLY FIRST TIME
	stockSQL := ""
	if isGenesis {
		stockSQL = fmt.Sprintf("SELECT ID, Name, GroupID, Cost, Price, StockQty, StockValue, LastBuyDate, LastsellDate, EditDate FROM fss.dbo.bsItem WHERE GroupID IN ('C', 'C-1', 'E') ORDER BY EditDate DESC;")
	} else {
	// STOCK sql STATMENT
		stockSQL = fmt.Sprintf("SELECT ID, Name, GroupID, Cost, Price, StockQty, StockValue, LastBuyDate, LastsellDate, EditDate FROM fss.dbo.bsItem WHERE EditDate >= '%s 00:00:00' AND GroupID IN ('C', 'C-1', 'E') ORDER BY EditDate DESC;", datetime)
	}
	
	
	// EXECUTE QUERY
	rows, err := db.QueryContext(ctx, stockSQL)
	if err != nil {
		return store, err
	}

	// CLOSE CONNECTIOn
	defer rows.Close()

	// ITERATE THROUGH THE RESULT SET.
	for rows.Next() {
		var stockRow models.Stock

		// GET VALUES FROM ROW.
		err := rows.Scan(
			&stockRow.ID, &stockRow.Name, &stockRow.GroupID,
			&stockRow.Cost, &stockRow.Price, &stockRow.StockQty,
			&stockRow.StockValue, &stockRow.LastBuyDate, &stockRow.LastSellDate,
			&stockRow.EditDate,
		)

		if err != nil {
			return store, err
		}

		store[stockRow.ID] = stockRow

		// IF YOU WANT TO SEE DATA STREAM FROM LOCAL DATABASE
		// fmt.Printf("ID: %s, Name: %s, GroupID: %s, Cost: %s, Price: %s, EditDate: %s\n",
		// 	stockRow.ID, stockRow.Name, stockRow.GroupID, stockRow.Cost, stockRow.Price, stockRow.EditDate)
	}

	return store, nil
}
