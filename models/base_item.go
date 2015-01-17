package models

import (
	"fmt"
	"github.com/fhbzyc/c_game/libs/db"
	"github.com/fhbzyc/c_game/models/table"
)

var BaseItem baseItemModel

type baseItemModel struct {
}

func (this baseItemModel) FindAll() ([]table.BaseItemTable, error) {
	var list []table.BaseItemTable
	return list, db.DataBase.Find(&list)
}

func (this baseItemModel) FindOne(itemId int) table.BaseItemTable {
	list, _ := this.FindAll()

	for _, item := range list {
		if item.ItemId == itemId {
			return item
		}
	}

	panic(fmt.Sprintf("不存在的道具: %d", itemId))
}
