package config

import (
	"github.com/wu-xing/wood-serve/common"
	"github.com/wu-xing/wood-serve/database"
)

func Migration(db *database.DB) {

	var count int
	err := common.GetDB().Connection.DB().QueryRow("select count(*) from app_config").Scan(&count)

	if err != nil {
		panic(err)
	}

	if count >= 1 {
		return
	}

	var appConfig = AppConfig{
		AppConfigData: &AppConfigData{
			NoteTypeConfigs: []NoteTypeProp{
				{
					Name:         "Markdown",
					CanExportPdf: false,
				},
			},
		},
	}

	err = common.GetDB().Connection.Create(appConfig).Error

	if err != nil {
		panic(err)
	}
}
