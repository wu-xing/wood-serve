package entitys

import (
	"time"
)

type ArticleBox struct {
	ID        string `gorm:"type:uuid; primary_key; default: uuid_generate_v4();"`
	Title     string `gorm:"size:100"`
	Creator   User   `gorm:"not null"`
	CreatorId string `gorm:"type:varchar(100)"`
	Type      string `gorm:"size:30;default:'DEFAULT'"`
	Status    string `gorm:"size:30;default:'ACTIVE'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
