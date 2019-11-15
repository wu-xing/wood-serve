package config

import (
	"database/sql/driver"
	"encoding/json"
)

type AppConfig struct {
	AppConfigData *AppConfigData `gorm:"type:JSON"`
}

func (AppConfig) TableName() string {
	return "app_config"
}

type AppConfigData struct {
	NoteTypeConfigs []NoteTypeProp `json:"noteTypeConfigs"`
}

func (acd *AppConfigData) Scan(value interface{}) error {
	if value == nil {
		*acd = AppConfigData{}
		return nil
	}
	t := AppConfigData{}
	if e := json.Unmarshal(value.([]byte), &t); e != nil {
		return e
	}
	*acd = t
	return nil
}

func (acd *AppConfigData) Value() (driver.Value, error) {
	if acd == nil {
		return nil, nil
	}
	b, e := json.Marshal(*acd)
	return b, e
}

type NoteTypeProp struct {
	Name         string `json:"name"`
	CanExportPdf bool   `json:"canExportPdf"`
}
