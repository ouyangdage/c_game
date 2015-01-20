package table

import (
	"github.com/fhbzyc/c_game/libs/db"
)

func init() {
	db.DataBase.Sync2(new(BaseTaskTable))
}

const (
	_ TaskType = iota
	TASK_LEVEL
	TASK_HERO_CLASS
)

type (
	TaskType int

	BaseTaskTable struct {
		TaskId   int      `xorm:"pk"`
		Type     TaskType `xorm:"'task_type' NOT NULL DEFAULT 0 TINYINT(2)"`
		Name     string   `xorm:"'task_name' NOT NULL DEFAULT ''"`
		Desc     string   `xorm:"'task_desc' NOT NULL TEXT"`
		Reward   string   `xorm:"'task_reward' NOT NULL TEXT"`
		PreId    int      `xorm:"'task_pre_id' NOT NULL DEFAULT 0"`
		MinLevel int      `xorm:"'task_min_level' NOT NULL DEFAULT 0"`
		Request  int      `xorm:"'task_request' NOT NULL DEFAULT 0"`
		Goal     int      `xorm:"'task_goal' NOT NULL DEFAULT 0 TINYINT(2)"`
	}
)

func (this BaseTaskTable) TableName() string {
	return "base_task"
}
