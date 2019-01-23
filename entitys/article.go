package entitys

import (
	"time"
)

type Article struct {
	ID           string      `gorm:"type:uuid; primary_key; default: uuid_generate_v4();" json:"id"`
	Title        string      `gorm:"size:350" json:"title"`
	Creator      *User       `gorm:"not null" json:"-"`
	CreatorId    string      `gorm:"type:varchar(100)" json:"-"`
	Box          *ArticleBox `json:"-"`
	BoxId        string      `gorm:"type:varchar(100)" json:"-"`
	Content      string      `gorm:"type:text" json:"content"`
	Type         string      `gorm:"size:30;default:'DEFAULT'" json:"type"`
	Status       string      `gorm:"size:30;default:'ACTIVE'" json:"status"`
	IsEncryption bool        `gorm:"default:false" json:"isEncryption"` // set num to auto incrementable
	IsPublic     bool        `gorm:"default:false" json:"isPublic"`     // create index with name `addr` for address
	CreatedAt    time.Time   `json:"createdAt"`
	UpdatedAt    time.Time   `json:"updatedAt"`
	DeletedAt    *time.Time
}
