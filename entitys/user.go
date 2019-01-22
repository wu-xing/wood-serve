package entitys

import (
	"github.com/satori/go.uuid"
	"time"
)

type User struct {
	ID        *uuid.UUID `gorm:"not null; primary_key; type:uuid; default: uuid_generate_v4();"`
	Username  string     `gorm:"type:varchar(100);unique_index"`
	Hash      string     `gorm:"size:255;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
