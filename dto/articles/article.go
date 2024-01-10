package dto

type ArticleRes struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Image       string `json:"image"`
	Description string `json:"Description"`
	UserID      int    `json:"user_id"`
}
