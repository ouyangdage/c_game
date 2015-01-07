package table

import (
	"github.com/fhbzyc/c_game/libs/db"
)

func init() {
	db.DataBase.Sync(new(BaseStoryTable))
}

type BaseStoryTable struct {
	StoryId int    `xorm:"pk"`
	Type    int    `xorm:"'story_type' pk TINYINT(1)"`
	Name    string `xorm:"'story_name' NOT NULL DEFAULT '' VARCHAR(255)"`
	Coin    string `xorm:"'story_coin' NOT NULL DEFAULT 0 INT(11)"`
}

func (this BaseStoryTable) TableName() string {
	return "base_story"
}
