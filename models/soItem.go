package models

import "database/sql"

type SOItem struct {
	RowOrder     string         `db:"RowOrder"`
	DocNo        string         `db:"DocNo"`
	SequenNo     string         `db:"SequenNo"`
	ItemID       string         `db:"ItemID"`
	ItemName     string         `db:"ItemName"`
	UnitID       string         `db:"UnitID"`
	StockID      string         `db:"StockID"`
	Price        string         `db:"Price"`
	Cost         sql.NullString `db:"Cost"`
	Qty          string         `db:"Qty"`
	RQty         string         `db:"RQty"`
	RID          string         `db:"RID"`
	RDate        sql.NullString `db:"RDate"`
	DlvDate      sql.NullString `db:"DlvDate"`
	CQty         string         `db:"CQty"`
	CID          string         `db:"CID"`
	CDate        sql.NullString `db:"CDate"`
	ApID         sql.NullString `db:"ApID"`
	TrackDate    sql.NullString `db:"TrackDate"`
	DiscountText string         `db:"DiscountText"`
	DiscountAmt  string         `db:"DiscountAmt"`
	TotalAmt     string         `db:"TotalAmt"`
	Description  sql.NullString `db:"Description"`
	AddID        string         `db:"AddID"`
	AddDate      sql.NullString `db:"AddDate"`
	EditID       string         `db:"EditID"`
	EditDate     sql.NullString `db:"EditDate"`
}
