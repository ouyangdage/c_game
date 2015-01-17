package models

import (
	"fmt"
	"github.com/fhbzyc/c_game/libs/db"
	"github.com/fhbzyc/c_game/models/table"
)

var BaseTask BaseTaskModel

type BaseTaskModel struct {
}

func (this BaseTaskModel) FindAll() ([]table.BaseTaskTable, error) {
	var list []table.BaseTaskTable
	return list, db.DataBase.Find(&list)
}

func (this BaseTaskModel) FindOne(id int) table.BaseTaskTable {
	list, _ := this.FindAll()

	for _, item := range list {
		if item.TaskId == id {
			return item
		}
	}

	panic(fmt.Sprintf("不存在的任务: %d", id))
}
