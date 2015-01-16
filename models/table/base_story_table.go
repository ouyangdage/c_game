package table

import (
	"github.com/fhbzyc/c_game/libs/db"
)

func init() {
	db.DataBase.Sync(new(BaseStoryTable))
}

type BaseStoryTable struct {
	StoryId      int    `xorm:"pk"`
	Type         int    `xorm:"'story_type' pk TINYINT(1)"`
	Name         string `xorm:"'story_name' NOT NULL DEFAULT ''"`
	Coin         int    `xorm:"'story_coin' NOT NULL DEFAULT 0"`
	HeroExp      int    `xorm:"'story_hero_exp' NOT NULL DEFAULT 0"`
	RoleExp      int    `xorm:"NOT NULL DEFAULT 0 INT(11)"`
	UnlockLevel  int    `xorm:"'story_unlock_level' NOT NULL DEFAULT 0"`
	Items        string `xorm:"'story_items' NOT NULL DEFAULT ''"`
	WipeOutItems string `xorm:"'story_wipe_out_items' NOT NULL DEFAULT ''"`
}

func (this BaseStoryTable) TableName() string {
	return "base_story"
}
