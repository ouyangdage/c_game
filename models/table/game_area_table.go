package table

import (
	"github.com/fhbzyc/c_game/libs/db"
)

func init() {
	db.DataBase.Sync(new(GameAreaTable))
}

type GameAreaTable struct {
	AreaId int    `xorm:"pk autoincr"`
	Name   string `xorm:"'area_name' NOT NULL DEFAULT '' VARCHAR(255)"`
}

func (this GameAreaTable) TableName() string {
	return "game_area"
}
