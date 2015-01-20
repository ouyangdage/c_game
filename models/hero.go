package models

import (
	"errors"
	"github.com/fhbzyc/c_game/libs/db"
	"github.com/fhbzyc/c_game/models/table"
)

var Hero heroModel = heroModel{
	MAXLEVEL: []int{10, 10, 10, 10, 15, 15, 15, 15, 15, 15, 15, 16, 16, 17, 17, 18, 18, 19, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90},
}

type heroModel struct {
	MAXLEVEL []int
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

func (this heroModel) MaxExp(roleLevel int) int {
	maxLevel := this.MAXLEVEL[roleLevel-1]
	return table.HeroExpToLevel[maxLevel-1]
}

func (this heroModel) AddExp(hero *table.HeroTable, exp, roleExp int) error {
	maxExp := this.MaxExp(roleExp)
	if hero.Exp >= maxExp {
		return errors.New("经验已是最大值")
	} else {
		hero.Exp += exp
		if hero.Exp > maxExp {
			hero.Exp = maxExp
		}
		return this.Update(hero)
	}
}

func (this heroModel) SoulNum(heroId int) int {
	star := BaseHero.FindOne(heroId).Star
	if star == 2 {
		return 18
	} else if star == 3 {
		return 30
	} else {
		return 7
	}
}
