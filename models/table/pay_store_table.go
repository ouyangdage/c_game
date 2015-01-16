package table

import (
	"github.com/fhbzyc/c_game/libs/db"
)

func init() {
	db.DataBase.Sync(new(PayStoreTable))
}

type PayStoreTable struct {
	GoodsId int    `xorm:"pk"`
	Type    int    `xorm:"'goods_type' NOT NULL DEFAULT 0 TINYINT(1)"`
	Money   int    `xorm:"'goods_money' NOT NULL DEFAULT 0"`
	Gold    int    `xorm:"'goods_gold' NOT NULL DEFAULT 0"`
	AddGold int    `xorm:"'goods_add_gold' NOT NULL DEFAULT 0"`
	Name    string `xorm:"'goods_name' NOT NULL DEFAULT ''"`
	Desc    string `xorm:"'goods_desc' NOT NULL TEXT"`
}

func (this PayStoreTable) TableName() string {
	return "pay_store"
}
