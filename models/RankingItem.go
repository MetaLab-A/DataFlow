package models

type RankingItem struct {
	ItemID    string  `db:"ItemID"`
	ItemName  string  `db:"ItemName"`
	Price     float64 `db:"Price"`
	Cost      float64 `db:"Cost"`
	Qty       int     `db:"Qty"`
	TotalAmt  float64 `db:"TotalAmt"`
	Margin    float64 `db:"Margin"`
	ProfitAmt float64 `db:"ProfitAmt"`
}
