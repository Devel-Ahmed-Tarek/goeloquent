package models

// User نموذج المستخدم مثال، يرث من BaseModel
type User struct {
	BaseModel
	Name   string `json:"name"`
	Email  string `json:"email" gorm:"unique"`
	Status string `json:"status"`
	Posts  []Post `json:"posts" gorm:"foreignKey:UserID"`
}
