package entitys

import (
	"time"
)

type ArticleBox struct {
	ID        string `gorm:"type:uuid; primary_key; default: uuid_generate_v4();" json:"id"`
	Name      string `gorm:"size:100" json:"name"`
	Creator   User
	CreatorId string     `gorm:"type:varchar(100)" json:"-"`
	Type      string     `gorm:"size:30;default:'DEFAULT'" json:"type"`
	Status    string     `gorm:"size:30;default:'ACTIVE'" json:"status"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
