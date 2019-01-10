// CREATE TABLE IF NOT EXISTS users(
//         id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
//         username TEXT NOT NULL UNIQUE,
//         hash TEXT NOT NULL,
//         created_at DATE,
//         updated_at DATE
// );

package entitys

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	ID           string  `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Username     string  `gorm:"type:varchar(100);unique_index"`
	Hash         string  `gorm:"size:255"`        // set field size to 255
	MemberNumber *string `gorm:"unique;not null"` // set member number to unique and not null
	Num          int     `gorm:"AUTO_INCREMENT"`  // set num to auto incrementable
	Address      string  `gorm:"index:addr"`      // create index with name `addr` for address
}
