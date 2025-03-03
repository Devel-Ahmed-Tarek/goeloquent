package models

// Post نموذج المشاركة مع علاقة BelongsTo مع User
type Post struct {
	BaseModel
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID uint   `json:"user_id"`
}
