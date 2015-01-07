package models

import (
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
