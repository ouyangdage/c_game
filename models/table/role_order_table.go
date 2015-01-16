package table

import (
	"github.com/fhbzyc/c_game/libs/db"
)

func init() {
	db.DataBase.Sync(new(OrderTable))
}

type OrderTable struct {
	OrderId int    `xorm:"pk autoincr"`
	RoleId  int    `xorm:"NOT NULL DEFAULT 0"`
	Type    int    `xorm:"'order_type' NOT NULL DEFAULT 0 TINYINT(1)"`
	Money   int    `xorm:"'order_money' NOT NULL DEFAULT 0"`
	Gold    int    `xorm:"'order_gold' NOT NULL DEFAULT 0"`
	AddGold int    `xorm:"'order_add_gold' NOT NULL DEFAULT 0"`
	Time    string `xorm:"'order_time' NOT NULL DEFAULT CURRENT_TIMESTAMP TIMESTAMP"`
	Status  int    `xorm:"'order_status' NOT NULL DEFAULT 0 TINYINT(1)"`
}

func (this OrderTable) TableName() string {
	return "role_order"
}
