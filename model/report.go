package model

type Report struct {
	TotalOrder int    `json:"total_order"`
	Month      string `json:"month"`
}

type CustomerReport struct {
	CustomerName string `json:"customer_name"`
	TotalOrder   int    `json:"total_order"`
	Month        string `json:"month"`
}

type AreaReport struct {
	AreaName   string `json:"area_name"`
	City       string `json:"city"`
	TotalOrder int    `json:"total_order"`
}

type HourReport struct {
	Hour       int    `json:"hour"`
	TimeRange  string `json:"time_range"`
	TotalOrder int    `json:"total_order"`
	Category   string `json:"category"`
}
