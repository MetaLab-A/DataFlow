package localsql

import (
	"context"
	"database/sql"
	"fmt"
)

// Stock data table from db
type Stock struct {
	ID       string
	Name     string
	GroupID  string
	Cost     []string
	Price    []string
	EditDate string
}


// ReadStockSQL reads all Stock from bsItem records
func ReadStockSQL(db *sql.DB, datetime string) (map[string]Stock, error) {
	store := make(map[string]Stock)

	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		return store, err
	}

	// Get current date
	// Stock SQL statment
	stockSQL := fmt.Sprintf("SELECT ID, Name, GroupID, Cost, Price, EditDate FROM fss.dbo.bsItem WHERE EditDate >= '%s 00:00:00' AND GroupID IN ('C', 'C-1', 'E') ORDER BY EditDate DESC;", datetime)

	// Execute query
	rows, err := db.QueryContext(ctx, stockSQL)
	if err != nil {
		return store, err
	}

	// Close connection
	defer rows.Close()

	// Iterate through the result set.
	for rows.Next() {
		var stockRow Stock

		// Get values from row.
		cost := ""
		price := ""
		err := rows.Scan(
			&stockRow.ID, &stockRow.Name, &stockRow.GroupID,
			&cost, &price, &stockRow.EditDate,
		)

		if err != nil {
			return store, err
		}

		stockRow.Cost = []string{cost}
		stockRow.Price = []string{price}
		store[stockRow.ID] = stockRow

		fmt.Printf("ID: %s, Name: %s, GroupID: %s, Cost: %s, Price: %s, EditDate: %s\n",
			stockRow.ID, stockRow.Name, stockRow.GroupID, stockRow.Cost, stockRow.Price, stockRow.EditDate)
	}

	return store, nil
}
