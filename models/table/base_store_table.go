package table

import (
	"github.com/fhbzyc/c_game/libs/db"
)

func init() {
	db.DataBase.Sync(new(BaseStoreTable))
}

type StoreType int

const (
	_ StoreType = iota
	STORE_TYPE_NORMAL

	GOODS_MONEY_COIN int = 1
	GOODS_MONEY_GOLD int = 2
)

type BaseStoreTable struct {
	GoodsId   int    `xorm:"pk autoincr"`
	ItemId    int    `xorm:"NOT NULL DEFAULT 0"`
	Type      int    `xorm:"'goods_type' pk TINYINT(1)"`
	Name      string `xorm:"'goods_name' NOT NULL DEFAULT ''"`
	Price     int    `xorm:"'goods_price' NOT NULL DEFAULT 0"`
	Num       int    `xorm:"'goods_num' NOT NULL DEFAULT 0"`
	Desc      string `xorm:"'goods_desc' NOT NULL TEXT"`
	StoreType int    `xorm:"NOT NULL DEFAULT 0"`
	Discount  int    `xorm:"'goods_discount' NOT NULL DEFAULT 0"`
	Time      int64  `xorm:"-"`
}

func (this BaseStoreTable) TableName() string {
	return "base_store"
}
