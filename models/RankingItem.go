package models

type RankingItem struct {
	ItemID     string  `db:"ItemID"`
	ItemName   string  `db:"ItemName"`
	highPrice  float64 `db:"highPrice"`
	lowPrice   float64 `db:"lowPrice"`
	highCost   float64 `db:"highCost"`
	lowCost    float64 `db:"lowCost"`
	highMargin float64 `db:"highMargin"`
	lowMargin  float64 `db:"lowMargin"`
	Qty        int     `db:"Qty"`
	TotalAmt   float64 `db:"TotalAmt"`
	ProfitAmt float64 `db:"ProfitAmt"`
}
