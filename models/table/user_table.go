package table

import (
	"github.com/fhbzyc/c_game/libs/db"
)

func init() {
	db.DataBase.Sync2(new(UserTable))
	// DataBase().MapCacher(new(UserTable), xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000))
}

// user
type UserTable struct {
	Uid          int  `xorm:"NOT NULL pk autoincr INT(11)"`
	PlatformId   int  `xorm:"NOT NULL DEFAULT 0 unique(PLAT_ID) TINYINT(2)"`
	PlatformUuid string `xorm:"NOT NULL DEFAULT '' unique(PLAT_ID) VARCHAR(100)"`
	Imei         string `xorm:"NOT NULL DEFAULT '' VARCHAR(50)"`
	Ip           string `xorm:"NOT NULL DEFAULT '' VARCHAR(15)"`
}

func (this UserTable) TableName() string {
	return "user"
}
