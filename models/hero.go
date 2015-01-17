package models

import (
	"github.com/fhbzyc/c_game/libs/db"
	"github.com/fhbzyc/c_game/models/table"
)

var Hero heroModel

type heroModel struct {
}

func (this heroModel) FindAll(roleId int) []table.HeroTable {
	var list []table.HeroTable
	err := db.DataBase.Where("role_id = ?", roleId).Find(&list)
	if err != nil {
		panic(err.Error())
	}
	return list

}

func (this heroModel) Insert(hero *table.HeroTable) error {
	_, err := db.DataBase.Insert(hero)
	return err
}

func (this heroModel) Update(hero *table.HeroTable) error {
	_, err := db.DataBase.Update(hero)
	return err
}

func (this heroModel) FindOne(roleId, heroId int) (*table.HeroTable, error) {
	hero := new(table.HeroTable)
	find, err := db.DataBase.Where("role_id = ? AND hero_id = ?", roleId, heroId).Get(hero)
	if !find {
		return nil, err
	}
	return hero, err
}
