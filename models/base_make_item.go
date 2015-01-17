package models

import (
	"github.com/fhbzyc/c_game/libs/db"
	"github.com/fhbzyc/c_game/models/table"
)

var BaseMakeItem baseMakeItemModel

type baseMakeItemModel struct {
}

func (this baseMakeItemModel) FindAll() ([]table.BaseMakeItemTable, error) {
	var list []table.BaseMakeItemTable
	return list, db.DataBase.Find(&list)
}

func (this baseMakeItemModel) FindMaterial(itemId int) []table.BaseMakeItemTable {
	list, _ := this.FindAll()

	var result []table.BaseMakeItemTable

	for _, item := range list {
		if item.ItemId == itemId {
			result = append(result, item)
		}
	}

	return result
}
