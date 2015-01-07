package table

import (
	"github.com/fhbzyc/c_game/libs/db"
)

func init() {
	db.DataBase.Sync(new(BaseItemTable))
}

type BaseItemTable struct {
	ItemId int    `xorm:"pk autoincr"`
	Name   string `xorm:"'item_name' NOT NULL DEFAULT '' VARCHAR(255)"`
	Type   int    `xorm:"'item_type' NOT NULL DEFAULT 0 TINYINT(2)"`
	Coin   int    `xorm:"'item_sell_coin' NOT NULL DEFAULT 0 INT(11)"`
}

func (this BaseItemTable) TableName() string {
	return "base_item"
}
