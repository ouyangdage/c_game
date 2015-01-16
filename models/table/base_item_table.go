package table

import (
	"github.com/fhbzyc/c_game/libs/db"
)

func init() {
	db.DataBase.Sync(new(BaseItemTable))
}

type BaseItemTable struct {
	ItemId   int    `xorm:"pk autoincr"`
	Name     string `xorm:"'item_name' NOT NULL DEFAULT ''"`
	Type     int    `xorm:"'item_type' NOT NULL DEFAULT 0 TINYINT(2)"`
	Coin     int    `xorm:"'item_sell_coin' NOT NULL DEFAULT 0 "`
	MakeCoin int    `xorm:"'item_make_coin' NOT NULL DEFAULT 0"`
}

func (this BaseItemTable) TableName() string {
	return "base_item"
}
