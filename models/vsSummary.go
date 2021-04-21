package models

type VSSummary struct {
	ItemID   string `db:"ItemID"`
	Qty      string `db:"Qty"`
	TotalAmt string `db:"TotalAmt"`
}
