package table

import (
	"github.com/fhbzyc/c_game/libs/db"
)

func init() {
	db.DataBase.Sync(new(TaskTable))
}

type TaskTable struct {
	RoleId   int      `xorm:"pk"`
	TaskId   int      `xorm:"pk"`
	Type     TaskType `xorm:"'task_type' NOT NULL DEFAULT 0"`
	Num      int      `xorm:"'task_num' NOT NULL DEFAULT 0"`
	Complete bool     `xorm:"'task_is_complete' NOT NULL DEFAULT 0 TINYINT(1)"`
}

func (t TaskTable) TableName() string {
	return "role_task"
}
