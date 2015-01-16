package table

import (
	"github.com/fhbzyc/c_game/libs/db"
)

func init() {
	db.DataBase.Sync(new(BaseHeroEquipTable))
}

type BaseHeroEquipTable struct {
	HeroId int `xorm:"pk"`
	Class  int `xorm:"'hero_class' pk NOT NULL DEFAULT 0 TINYINT(1)"`
	Equip1 int `xorm:"'equip_id1' NOT NULL DEFAULT 0"`
	Equip2 int `xorm:"'equip_id2' NOT NULL DEFAULT 0"`
	Equip3 int `xorm:"'equip_id3' NOT NULL DEFAULT 0"`
	Equip4 int `xorm:"'equip_id4' NOT NULL DEFAULT 0"`
	Equip5 int `xorm:"'equip_id5' NOT NULL DEFAULT 0"`
	Equip6 int `xorm:"'equip_id6' NOT NULL DEFAULT 0"`
}

func (this BaseHeroEquipTable) TableName() string {
	return "base_hero_equip"
}
