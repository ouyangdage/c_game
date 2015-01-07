package table

import (
	"github.com/fhbzyc/c_game/libs/db"
)

func init() {
	db.DataBase.Sync(new(ItemTable))
}

type ItemTable struct {
	RoleId int `xorm:"pk"`
	ItemId int `xorm:"pk"`
	Num    int `xorm:"'item_num' NOT NULL DEFAULT 0 INT(11)"`
}

func (i ItemTable) TableName() string {
	return "role_item"
}
