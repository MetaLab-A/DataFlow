package models

import "database/sql"

type POItem struct {
	RowOrder     string         `db:"RowOrder"`
	DocNo        string         `db:"DocNo"`
	SequenNo     string         `db:"SequenNo"`
	ItemID       string         `db:"ItemID"`
	ItemName     string         `db:"ItemName"`
	UnitID       string         `db:"UnitID"`
	StockID      string         `db:"StockID"`
	Price        string         `db:"Price"`
	Qty          string         `db:"Qty"`
	DiscountText string         `db:"DiscountText"`
	DiscountAmt  string         `db:"DiscountAmt"`
	TotalAmt     string         `db:"TotalAmt"`
	Description  sql.NullString `db:"Description"`
	AddID        string         `db:"AddID"`
	AddDate      sql.NullString `db:"AddDate"`
	EditID       string         `db:"EditID"`
	EditDate     sql.NullString `db:"EditDate"`
}
