package models

import (
	"github.com/fhbzyc/c_game/libs/db"
	"github.com/fhbzyc/c_game/models/table"
)

var Item ItemModel

type ItemModel struct {
}

func (this ItemModel) FindAll(roleId int) ([]table.ItemTable, error) {
	var list []table.ItemTable
	return list, db.DataBase.Where("role_id = ?", roleId).Find(&list)
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

func (this ItemModel) Add(roleId, itemId, num int) (*table.ItemTable, error) {

	item := new(table.ItemTable)

	_, err := db.DataBase.Where("role_id = ? AND item_id = ?", roleId, itemId).Get(item)
	if err != nil {
		return item, err
	}

	if item.ItemId == 0 {
		item.ItemId = itemId
		item.RoleId = roleId
	} else {
		item.Num += num
	}

	_, err = db.DataBase.Insert(item)
	return item, err
}