package table

import (
	"github.com/fhbzyc/c_game/libs/db"
)

func init() {
	db.DataBase.Sync(new(ShopTable))
}

type ShopTable struct {
	RoleId int  `xorm:"pk"`
	Goods1 int  `xorm:"NOT NULL DEFAULT 0"`
	Goods2 int  `xorm:"NOT NULL DEFAULT 0"`
	Goods3 int  `xorm:"NOT NULL DEFAULT 0"`
	Goods4 int  `xorm:"NOT NULL DEFAULT 0"`
	Goods5 int  `xorm:"NOT NULL DEFAULT 0"`
	Goods6 int  `xorm:"NOT NULL DEFAULT 0"`
	Goods7 int  `xorm:"NOT NULL DEFAULT 0"`
	Goods8 int  `xorm:"NOT NULL DEFAULT 0"`
	IsBuy1 bool `xorm:"NOT NULL DEFAULT 0 TINYINT(1)"`
	IsBuy2 bool `xorm:"NOT NULL DEFAULT 0 TINYINT(1)"`
	IsBuy3 bool `xorm:"NOT NULL DEFAULT 0 TINYINT(1)"`
	IsBuy4 bool `xorm:"NOT NULL DEFAULT 0 TINYINT(1)"`
	IsBuy5 bool `xorm:"NOT NULL DEFAULT 0 TINYINT(1)"`
	IsBuy6 bool `xorm:"NOT NULL DEFAULT 0 TINYINT(1)"`
	IsBuy7 bool `xorm:"NOT NULL DEFAULT 0 TINYINT(1)"`
	IsBuy8 bool `xorm:"NOT NULL DEFAULT 0 TINYINT(1)"`
	Date   int  `xorm:"'shop_date' NOT NULL DEFAULT 0"`
}

func (this ShopTable) TableName() string {
	return "role_shop"
}
