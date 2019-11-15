package migration

import (
	"github.com/wu-xing/wood-serve/common"
	"github.com/wu-xing/wood-serve/domain/config"
)

func AppMigration() {
	db := common.GetDB()

	config.Migration(db)
}
