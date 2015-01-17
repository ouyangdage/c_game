package models

import (
	"fmt"
	"github.com/fhbzyc/c_game/libs/db"
	"github.com/fhbzyc/c_game/models/table"
)

var BaseHeroEquip baseHeroEquip

type baseHeroEquip struct {
}

func (this baseHeroEquip) FindAll() ([]table.BaseHeroEquipTable, error) {
	var list []table.BaseHeroEquipTable
	return list, db.DataBase.Find(&list)
}

func (this baseHeroEquip) FindOne(heroId, class int) *table.BaseHeroEquipTable {

	list, _ := this.FindAll()

	for _, item := range list {
		if item.HeroId == heroId && item.Class == class {
			return &item
		}
	}

	panic(fmt.Sprintf("不存在的阶级: %d , %d", heroId, class))
}
