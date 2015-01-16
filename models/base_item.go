package models

import (
	"fmt"
	"github.com/fhbzyc/c_game/libs/db"
	"github.com/fhbzyc/c_game/models/table"
)

var BaseItem BaseItemModel

type BaseItemModel struct {
}

func (this BaseItemModel) FindAll() ([]table.BaseItemTable, error) {
	var list []table.BaseItemTable
	return list, db.DataBase.Find(&list)
}

func (this BaseItemModel) FindOne(itemId int) table.BaseItemTable {
	list, _ := this.FindAll()

	for _, item := range list {
		if item.ItemId == itemId {
			return item
		}
	}

	panic(fmt.Sprintf("不存在的道具: %d", itemId))
}
