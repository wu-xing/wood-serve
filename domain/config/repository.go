package config

import (
	"github.com/jinzhu/gorm"
)

type Repository interface {
	GetConfig() *AppConfig
}

type RepositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{db: db}
}

func (r *RepositoryImpl) GetConfig() *AppConfig {
	appConfig := AppConfig{}
	r.db.Last(&appConfig)
	return &appConfig
}
