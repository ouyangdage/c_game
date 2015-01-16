package models

import (
	"fmt"
	"github.com/fhbzyc/c_game/libs/db"
	"github.com/fhbzyc/c_game/models/table"
)

var BaseStore BaseStoreModel

type BaseStoreModel struct {
}

func (this BaseStoreModel) FindAll() ([]table.BaseStoreTable, error) {
	var list []table.BaseStoreTable
	return list, db.DataBase.Find(&list)
}

func (this BaseStoreModel) FindOne(id int) table.BaseStoreTable {
	list, _ := this.FindAll()

	for _, item := range list {
		if item.GoodsId == id {
			return item
		}
	}

	panic(fmt.Sprintf("不存在的Id: %d", id))
}
