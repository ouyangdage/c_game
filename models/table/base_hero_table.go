package table

import (
	"github.com/fhbzyc/c_game/libs/db"
)

func init() {
	db.DataBase.Sync(new(BaseHeroTable))
}

type BaseHeroTable struct {
	HeroId int    `xorm:"pk autoincr"`
	Name   string `xorm:"'hero_name' NOT NULL DEFAULT '' VARCHAR(255)"`
	Star   int    `xorm:"'hero_star' NOT NULL DEFAULT 0 TINYINT(1)"`
	Skill1 int    `xorm:"'skill_id_1' NOT NULL DEFAULT 0 INT(11)"`
	Skill2 int    `xorm:"'skill_id_2' NOT NULL DEFAULT 0 INT(11)"`
	Skill3 int    `xorm:"'skill_id_3' NOT NULL DEFAULT 0 INT(11)"`
	Skill4 int    `xorm:"'skill_id_4' NOT NULL DEFAULT 0 INT(11)"`
}

func (this BaseHeroTable) TableName() string {
	return "base_hero"
}
