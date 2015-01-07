package table

import (
	"github.com/fhbzyc/c_game/libs/db"
)

func init() {
	db.DataBase.Sync(new(ServerTable))
}

type PlatformTable struct {
	PlatformId   int    `xorm:"NOT NULL pk autoincr INT(11)"`
	PlatformName string `xorm:"NOT NULL DEFAULT '' VARCHAR(255)"`
}
