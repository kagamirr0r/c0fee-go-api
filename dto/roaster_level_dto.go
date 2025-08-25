package dto

type RoastLevelOutput struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type RoastLevelsOutput struct {
	RoastLevels []RoastLevelOutput `json:"roast_levels"`
	Count       uint               `json:"count"`
}
