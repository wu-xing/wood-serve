package entitys

import (
	"time"
)

type Image struct {
	ID        string `gorm:"type:uuid; primary_key"`
	CreatorId string `gorm:"type:varchar(100)"`
	Filename  string `gorm:"size:100;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
