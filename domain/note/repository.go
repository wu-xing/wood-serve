package note

import (
	"github.com/jinzhu/gorm"
)

type Repository interface {
	GetAllNotes() *[]Note
}

type RepositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{db: db}
}

func (r *RepositoryImpl) GetAllNotes(id string) *[]Note {
	notes := []Note{}
	r.db.Find(notes)
	return &notes
}
