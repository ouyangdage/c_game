package table

import (
	"github.com/fhbzyc/c_game/libs/db"
)

func init() {
	db.DataBase.Sync(new(BaseSkillTable))
}

type BaseSkillTable struct {
	SkillId int    `xorm:"pk autoincr"`
	Name    string `xorm:"'skill_name' NOT NULL DEFAULT '' VARCHAR(255)"`
}

func (s BaseSkillTable) TableName() string {
	return "base_skill"
}
