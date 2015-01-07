package models

import (
	"github.com/fhbzyc/c_game/libs/db"
	"github.com/fhbzyc/c_game/models/table"
)

var Hero HeroModel

type HeroModel struct {
}

func (this HeroModel) FindAll(roleId int) ([]table.HeroTable, error) {
	var result []table.HeroTable
	return result, db.DataBase.Where("role_id = ?", roleId).Find(&result)
}

func (this HeroModel) Insert(hero *table.HeroTable) error {
	_, err := db.DataBase.Insert(hero)
	return err
}

func (this HeroModel) Update(hero *table.HeroTable) error {
	_, err := db.DataBase.Update(hero)
	return err
}
