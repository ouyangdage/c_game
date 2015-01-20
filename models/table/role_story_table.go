package table

import (
	"github.com/fhbzyc/c_game/libs/db"
)

func init() {
	db.DataBase.Sync2(new(StoryTable))
}

type StoryTable struct {
	RoleId  int       `xorm:"pk"`
	StoryId int       `xorm:"pk"`
	Type    StoryType `xorm:"pk 'stroy_type' NOT NULL DEFAULT 0 TINYINT(1)"`
	Star    int       `xorm:"'stroy_star' NOT NULL DEFAULT 0 TINYINT(1)"`
	Num     int       `xorm:"'story_num' NOT NULL DEFAULT 0"`
	BuyNum  int       `xorm:"'story_buy_num' NOT NULL DEFAULT 0"`
	Date    int       `xorm:"'story_date' NOT NULL DEFAULT 0"`
}

func (s StoryTable) TableName() string {
	return "role_story"
}
