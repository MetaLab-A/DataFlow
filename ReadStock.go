package main

import (
	"database/sql"
	"context"
	"fmt"
	"time"
)


// ReadStock reads all Stock from bsItem records
func ReadStock(db *sql.DB) (int, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	date := time.Now().Format("2006-01-02")
	fmt.Println(date)
	stockSQL := fmt.Sprintf("SELECT ID, Name, GroupID, Cost, Price, EditDate FROM fss.dbo.bsItem WHERE EditDate >= '%s 00:00:00' ORDER BY EditDate DESC;", date)

	// Execute query
	rows, err := db.QueryContext(ctx, stockSQL)
	if err != nil {
		return -1, err
	}

	defer rows.Close()

	var count int

	// Iterate through the result set.
	for rows.Next() {
		var (
			id, name string
			groupID, cost, price string
			editDate string
		)

		// Get values from row.
		err := rows.Scan(
			&id, &name, &groupID, 
			&cost, &price, &editDate,
		)

		if err != nil {
			return -1, err
		}

		fmt.Printf("ID: %s, Name: %s, GroupID: %s, Cost: %s, Price: %s, EditDate: %s\n",
					id, name, groupID, cost, price, editDate)
		count++
	}

	return count, nil
}