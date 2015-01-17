package models

import (
	"github.com/fhbzyc/c_game/libs/db"
	"github.com/fhbzyc/c_game/libs/log"
	"github.com/fhbzyc/c_game/models/table"
	// "github.com/fhbzyc/c_game/network"
)

func init() {
	go Task.taskRun(taskChan)
}

var (
	Task     taskModel
	taskChan chan *taskParam = make(chan *taskParam, 1000)
)

func (this taskModel) FindAll(roleId int) []table.TaskTable {
	var list []table.TaskTable
	err := db.DataBase.Where("role_id = ?", roleId).Find(&list)
	if err != nil {
		panic(err.Error())
	}
	return list
}

func (this taskModel) FindByType(roleId int, typ table.TaskType) *table.TaskTable {
	tasks := this.FindAll(roleId)
	for _, task := range tasks {
		if task.Type == typ {
			return &task
		}
	}
	return nil
}

type taskModel struct {
}

type taskParam struct {
	*table.RoleTable
	Type   table.TaskType
	Target int
	Num    int
}

func (this taskModel) Add(typ *taskParam) {
	taskChan <- typ
}

func (this taskModel) taskRun(channel <-chan *taskParam) {

	defer func() {
		if err := recover(); err != nil {
			log.Logger.Error("task panic : %v", err)
			this.taskRun(channel)
		}
	}()

	for task := range channel {
		this.check(task)
	}
}

func (this taskModel) check(param *taskParam) {

	task := this.FindByType(param.RoleId, param.Type)
	if task != nil {

		upOK := false

		switch task.Type {
		case table.TASK_LEVEL:
			upOK = true
		case table.TASK_HERO_CLASS:
			upOK = true
		}

		find := false

		if upOK {
			if param.Target >= BaseTask.FindOne(task.TaskId).Request {
				find = true
			}
		} else {
			if param.Target == BaseTask.FindOne(task.TaskId).Request {
				find = true
			}
		}

		if find {
			task.Num += param.Num
			db.DataBase.Update(task)
			//network.SendMessage(role.AreaId, role.RoleId, []byte("任务有更新 转小红点"))
		}
	}

}

//func (this taskModel) find(task *table.TaskTable, num int) bool {
//	task.Num += num
//	db.DataBase.Update(task)
//	return true
//}

//func (this taskModel) checkLevel(task *table.TaskTable, level, num int) bool {

//	if level >= BaseTask.FindOne(task.TaskId).Request {
//		task.Num += num
//		db.DataBase.Update(task)
//	}

//	return false
//}

//func (this taskModel) checkHeroClass(task *table.TaskTable, class, num int) bool {

//	if class >= BaseTask.FindOne(task.TaskId).Request {
//		task.Num += num
//		db.DataBase.Update(task)
//	}

//	return false
//}

//func (this taskModel) checkStory(task *table.TaskTable, storyId, num int) bool {

//	if storyId == BaseTask.FindOne(task.TaskId).Request {
//		return this.find(task, num)
//	}

//	return false
//}

//// 挑战精英副本
//func (this taskModel) checkHeroStory(task *table.TaskTable, storyId, num int) bool {

//	if storyId == BaseTask.FindOne(task.TaskId).Request {
//		return this.find(task, num)
//	}

//	return false
//}
