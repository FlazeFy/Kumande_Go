package models

type (
	GetConsumeSearch struct {
		Slug           string `json:"slug_name"`
		ConsumeType    string `json:"consume_type"`
		ConsumeName    string `json:"consume_name"`
		ConsumeDetail  string `json:"consume_detail"`
		ConsumeFrom    string `json:"consume_from"`
		IsFavorite     string `json:"is_favorite"`
		ConsumeTag     string `json:"consume_tag"`
		ConsumeComment string `json:"consume_comment"`
		CreatedAt      string `json:"created_at"`
	}
)
