package table

import (
	"github.com/fhbzyc/c_game/libs/db"
)

func init() {
	db.DataBase.Sync(new(BaseTaskTable))
}

type BaseTaskTable struct {
	TaskId   int    `xorm:"pk"`
	Type     int    `xorm:"'task_type' NOT NULL DEFAULT 0 TINYINT(2)"`
	Name     string `xorm:"'task_name' NOT NULL DEFAULT '' VARCHAR(255)"`
	Desc     string `xorm:"'task_desc' NOT NULL DEFAULT '' VARCHAR(3000)"`
	Reward   string `xorm:"'task_reward' NOT NULL DEFAULT '' VARCHAR(3000)"`
	PreId    int    `xorm:"'task_pre_id' NOT NULL DEFAULT 0 INT(11)"`
	MinLevel int    `xorm:"'task_min_level' NOT NULL DEFAULT 0 INT(11)"`
	Enable   bool   `xorm:"'task_enable' NOT NULL DEFAULT 0 TINYINT(1)"`
}

func (this BaseTaskTable) TableName() string {
	return "base_task"
}
