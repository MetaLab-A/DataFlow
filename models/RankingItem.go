package models

type RankingItem struct {
	ItemID     string  `db:"ItemID"`
	ItemName   string  `db:"ItemName"`
	HighPrice  float64 `db:"HighPrice"`
	LowPrice   float64 `db:"LowPrice"`
	HighCost   float64 `db:"HighCost"`
	LowCost    float64 `db:"LowCost"`
	HighMargin float64 `db:"HighMargin"`
	LowMargin  float64 `db:"LowMargin"`
	Qty        int     `db:"Qty"`
	TotalAmt   float64 `db:"TotalAmt"`
	ProfitAmt  float64 `db:"ProfitAmt"`
}
