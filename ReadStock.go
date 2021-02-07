package main

import (
	"context"
	"database/sql"
	"fmt"
)

type Stock struct {
	id       string
	name     string
	groupID  string
	cost     string
	price    string
	editDate string
}

// ReadStock reads all Stock from bsItem records
func ReadStock(db *sql.DB, datetime string) ([]Stock, error) {
	store := []Stock{}

	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		return []Stock{}, err
	}

	// Get current date

	// Stock SQL statment
	stockSQL := fmt.Sprintf("SELECT ID, Name, GroupID, Cost, Price, EditDate FROM fss.dbo.bsItem WHERE EditDate >= '%s 00:00:00' AND GroupID IN ('C', 'C-1', 'E') ORDER BY EditDate DESC;", datetime)

	// Execute query
	rows, err := db.QueryContext(ctx, stockSQL)
	if err != nil {
		return []Stock{}, err
	}

	// Close connection
	defer rows.Close()

	// Iterate through the result set.
	for rows.Next() {
		var stockRow Stock

		// Get values from row.
		err := rows.Scan(
			&stockRow.id, &stockRow.name, &stockRow.groupID,
			&stockRow.cost, &stockRow.price, &stockRow.editDate,
		)

		if err != nil {
			return []Stock{}, err
		}

		store = append(store, stockRow)
		fmt.Printf("ID: %s, Name: %s, GroupID: %s, Cost: %s, Price: %s, EditDate: %s\n",
			stockRow.id, stockRow.name, stockRow.groupID, stockRow.cost, stockRow.price, stockRow.editDate)
	}

	return store, nil
}
