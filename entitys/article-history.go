package entitys

import (
	"time"
)

type ArticleHistory struct {
	ID        string    `gorm:"type:uuid; primary_key; default: uuid_generate_v4();" json:"id"`
	ArticleId string    `gorm:"size(36)" json:"articleId" json:"articleId"`
	Date      string    `gorm:"size:20" json:"date"`
	Content   string    `json:"content" json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
