package models

type RankingItem struct {
	ItemID       string `db:"ItemID"`
	ItemName     string `db:"ItemName"`
	Price        int64  `db:"Price"`
	Cost         int64  `db:"Cost"`
	Qty          int64  `db:"Qty"`
	DiscountText string `db:"DiscountText"`
	DiscountAmt  string `db:"DiscountAmt"`
	TotalAmt     string `db:"TotalAmt"`
	Margin       string `db:"Margin"`
	ProfitAmt    string `db:"ProfitAmt"`
}
