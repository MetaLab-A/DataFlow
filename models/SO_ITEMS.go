package models

import "database/sql"

type SO_ITEMS struct {
	DocNo			string	`db:"DocNo"`
	SequenNo		string	`db:"SequenNo"`
	ItemID			string	`db:"ItemID"`
	ItemName		string	`db:"ItemName"`
	UnitID			string	`db:"UnitID"`
	StockID			string	`db:"StockID"`
	Price			string	`db:"Price"`
	Cost			string	`db:"Cost"`
	Qty				string	`db:"QTY"`
	RQty			string	`db:"RQty"`
	RID				string	`db:"RID"`
	RDate			sql.NullString	`db:"RDate"`
	DlvDate			sql.NullString	`db:"DlvDate"`
	CQty			string	`db:"CQty"`
	CID				string	`db:"CID"`
	CDate			sql.NullString	`db:"CDate"`
	ApID			string	`db:"ApID"`
	TrackDate		sql.NullString	`db:"TrackDate"`
	DiscountText	string	`db:"DiscountText"`
	DiscountAmt		string	`db:"DiscountAmt"`
	TotalAmt		string	`db:"TotalAmt"`
	Description		string	`db:"Description"`
	AddID			string	`db:"AddID"`
	AddDate			sql.NullString	`db:"AddDate"`
	EditID			string	`db:"EditID"`
	EditDate		sql.NullString	`db:"EditDate"`
}