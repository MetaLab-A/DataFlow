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