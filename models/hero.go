package models

import (
	"github.com/fhbzyc/c_game/libs/db"
	"github.com/fhbzyc/c_game/models/table"
)

var Hero HeroModel

type HeroModel struct {
}

func (this HeroModel) FindAll(roleId int) []table.HeroTable {
	var list []table.HeroTable
	err := db.DataBase.Where("role_id = ?", roleId).Find(&list)
	if err != nil {
		panic(err.Error())
	}
	return list

}

func (this HeroModel) Insert(hero *table.HeroTable) error {
	_, err := db.DataBase.Insert(hero)
	return err
}

func (this HeroModel) Update(hero *table.HeroTable) error {
	_, err := db.DataBase.Update(hero)
	return err
}

func (this HeroModel) FindOne(roleId, heroId int) *table.HeroTable {
	hero := new(table.HeroTable)
	find, err := db.DataBase.Where("role_id = ? AND hero_id = ?", roleId, heroId).Get(hero)
	if err != nil {
		panic(err.Error())
	}
	if !find {
		return nil
	}
	return hero
}
