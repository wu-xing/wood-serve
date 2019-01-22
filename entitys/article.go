package entitys

import (
	"time"
)

type Article struct {
	ID           string `gorm:"type:uuid; primary_key; default: uuid_generate_v4();"`
	Title        string `gorm:"size:350"`
	Creator      *User  `gorm:"not null"`
	CreatorId    string `gorm:"type:varchar(100)"`
	Box          *ArticleBox
	BoxId        string `gorm:"type:varchar(100)"`
	Content      string `gorm:"type:text"`
	Type         string `gorm:"size:30;default:'DEFAULT'"`
	Status       string `gorm:"size:30;default:'ACTIVE'"`
	IsEncryption bool   `gorm:"default:false"` // set num to auto incrementable
	IsPublic     bool   `gorm:"default:false"` // create index with name `addr` for address
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}
