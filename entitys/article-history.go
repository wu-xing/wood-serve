package entitys

import (
	"time"
)

type ArticleHistory struct {
	ID        string `gorm:"type:uuid; primary_key"`
	ArticleId string `gorm:"size(36)" json:"articleId"`
	Date      string `gorm:"size:20"`
	Content   string `json:"content"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
