package models

import (
	"github.com/fhbzyc/c_game/libs/db"
	"github.com/fhbzyc/c_game/models/table"
)

var Item itemModel

type itemModel struct {
}

func (this itemModel) FindAll(roleId int) ([]table.ItemTable, error) {
	var list []table.ItemTable
	return list, db.DataBase.Where("role_id = ?", roleId).Find(&list)
}

func (this itemModel) FindOne(roleId, itemId int) (*table.ItemTable, error) {
	item := new(table.ItemTable)
	find, err := db.DataBase.Where("role_id = ? AND item_id = ?", roleId, itemId).Get(item)
	if !find {
		return nil, err
	}
	return item, err
}

func (this itemModel) Sub(item *table.ItemTable, num int) error {

	var err error
	item.Num -= num
	if item.Num <= 0 {
		_, err = db.DataBase.Delete(item)
	} else {
		_, err = db.DataBase.Update(item)
	}
	return err
}

func (this itemModel) Add(roleId, itemId, num int) *table.ItemTable {

	item, err := this.FindOne(roleId, itemId)
	if err != nil {
		panic(err)
	}

	if item == nil {
		item = new(table.ItemTable)
		item.ItemId = itemId
		item.RoleId = roleId
	}
	item.Num += num

	_, err = db.DataBase.Insert(item)
	if err != nil {
		panic(err)
	}

	return item
}
