package models

import (
	"github.com/fhbzyc/c_game/libs/db"
	"github.com/fhbzyc/c_game/models/table"
)

var Item ItemModel

type ItemModel struct {
}

func (this ItemModel) FindAll(roleId int) []table.ItemTable {
	var list []table.ItemTable
	err := db.DataBase.Where("role_id = ?", roleId).Find(&list)
	if err != nil {
		panic(err.Error())
	}
	return list
}

func (this ItemModel) FindOne(roleId, itemId int) *table.ItemTable {
	item := new(table.ItemTable)
	find, err := db.DataBase.Where("role_id = ? AND item_id = ?", roleId, itemId).Get(item)
	if err != nil {
		panic(err.Error())
	}
	if !find {
		return nil
	}
	return item
}

func (this ItemModel) Sub(item *table.ItemTable, num int) error {

	var err error
	item.Num -= num
	if item.Num <= 0 {
		_, err = db.DataBase.Delete(item)
	} else {
		_, err = db.DataBase.Update(item)
	}
	return err
}

func (this ItemModel) Add(roleId, itemId, num int) *table.ItemTable {

	item := this.FindOne(roleId, itemId)

	if item == nil {
		item = new(table.ItemTable)
		item.ItemId = itemId
		item.RoleId = roleId
	}
	item.Num += num

	_, err := db.DataBase.Insert(item)
	if err != nil {
		panic(err.Error())
	}

	return item
}
