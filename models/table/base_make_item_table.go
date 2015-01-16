package table

import (
	"github.com/fhbzyc/c_game/libs/db"
)

func init() {
	db.DataBase.Sync(new(BaseMakeItemTable))
}

type BaseMakeItemTable struct {
	ItemId     int `xorm:"pk"`
	MaterialId int `xorm:"pk"`
	Num        int `xorm:"'item_num' NOT NULL DEFAULT 0"`
}

func (s BaseMakeItemTable) TableName() string {
	return "base_make_item"
}
