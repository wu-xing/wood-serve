package entitys

import (
	"time"
)

type User struct {
	ID        string `gorm:"type:uuid; primary_key"`
	Username  string `gorm:"type:varchar(100);unique_index"`
	Hash      string `gorm:"size:255;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
