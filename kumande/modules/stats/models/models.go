package models

type (
	GetMostAppear struct {
		Context string `json:"context"`
		Total   int    `json:"total"`
	}
	GetSpendingInfo struct {
		TotalDays    int `json:"total_days"`
		TotalPayment int `json:"total_payment"`
	}
	GetBodyInfo struct {
		Weight    int    `json:"weight"`
		Height    int    `json:"height"`
		Result    int    `json:"result"`
		CreatedAt string `json:"created_at"`
	}
)
