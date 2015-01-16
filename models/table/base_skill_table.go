package table

import (
	"github.com/fhbzyc/c_game/libs/db"
)

func init() {
	db.DataBase.Sync(new(BaseSkillTable))
}

type BaseSkillTable struct {
	SkillId      int    `xorm:"pk autoincr"`
	Name         string `xorm:"'skill_name' NOT NULL DEFAULT '' VARCHAR(255)"`
	MaxLevel     int    `xorm:"'skill_max_level' NOT NULL DEFAULT 0 INT(11)"`
	BaseCoin     int    `xorm:"'skill_base_coin' NOT NULL DEFAULT 0 INT(11)"`
	Levelup_Coin int    `xorm:"'skill_levelup_coin' NOT NULL DEFAULT 0 INT(11)"`
}

func (s BaseSkillTable) TableName() string {
	return "base_skill"
}
