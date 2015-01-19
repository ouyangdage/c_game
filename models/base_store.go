package models

import (
	"github.com/fhbzyc/c_game/libs/db"
	"github.com/fhbzyc/c_game/models/table"
)

var BaseStore baseStoreModel

type baseStoreModel struct {
}

func (this baseStoreModel) FindAll(typ table.StoreType) ([]table.BaseStoreTable, error) {
	var list []table.BaseStoreTable
	return list, db.DataBase.Where("store_type = ?", typ).Find(&list)
}
