package table

import (
	"github.com/fhbzyc/c_game/libs/db"
)

func init() {
	db.DataBase.Sync(new(ServerTable))
}

type ServerTable struct {
	ServerId int  `xorm:"NOT NULL pk autoincr INT(11)"`
	ServerIp string `xorm:"NOT NULL DEFAULT '' VARCHAR(15)"`
}

func (this ServerTable) TableName() string {
	return "config_server"
}
