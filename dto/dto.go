package dto

// 汎用参照型
type IdRef struct {
	ID uint `json:"id"`
}

// 汎用 Output型
type IdNameSummary struct {
	ID   interface{} `json:"id"` // uint または string
	Name string      `json:"name"`
}
