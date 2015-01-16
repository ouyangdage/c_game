package table

import (
	"github.com/fhbzyc/c_game/libs/db"
)

func init() {
	db.DataBase.Sync(new(BaseStoreTable))
}

type BaseStoreTable struct {
	GoodsId int    `xorm:"pk autoincr"`
	Type    int    `xorm:"'goods_type' pk TINYINT(1)"`
	Name    string `xorm:"'goods_name' NOT NULL DEFAULT ''"`
	Price   int    `xorm:"'goods_price' NOT NULL DEFAULT 0"`
	Num     int    `xorm:"'goods_num' NOT NULL DEFAULT 0"`
	Desc    string `xorm:"'goods_desc' NOT NULL DEFAULT '' VARCHAR(3000)"`
}

func (this BaseStoreTable) TableName() string {
	return "base_store"
}
