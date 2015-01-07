package table

import (
	"github.com/fhbzyc/c_game/libs/db"
)

func init() {
	db.DataBase.Sync(new(HeroTable))
}

type HeroTable struct {
	HeroId int `xorm:"pk"`
	RoleId int `xorm:"pk"`
	Exp    int `xorm:"'hero_exp' NOT NULL DEFAULT 0 INT(11)"`
	Atk    int `xorm:"'hero_atk' NOT NULL DEFAULT 0 INT(11)"`
	Int    int `xorm:"'hero_int' NOT NULL DEFAULT 0 INT(11)"`
	Dex    int `xorm:"'hero_dex' NOT NULL DEFAULT 0 INT(11)"`
	Star   int `xorm:"'hero_star' NOT NULL DEFAULT 0 TINYINT(1)"`
	Class  int `xorm:"'hero_class' NOT NULL DEFAULT 0 TINYINT(2)"`
}

func (h HeroTable) TableName() string {
	return "role_hero"
}

//
//func (g *HeroTable) Insert() error {
//	_, err := DataBase().Insert(g)
//	return err
//}
//
//func (g *HeroTable) Update() error {
//	_, err := DataBase().Id(g.HeroId).Update(g)
//	return err
//}
//
//func (g *HeroTable) Delete() error {
//	_, err := DataBase().Id(g.HeroId).Delete(new(HeroTable))
//	return err
//}
//
//func (g *HeroTable) AddExp(exp int32) error {
//	g.HeroExp += exp
//	return g.Update()
//}
//
//func (g *HeroTable) AddAtk(atk int32) error {
//	g.HeroAtk += atk
//	return g.Update()
//}
//
//func (g *HeroTable) AddInt(int int32) error {
//	g.HeroInt += int
//	return g.Update()
//}
//
//func (g *HeroTable) AddDex(dex int32) error {
//	g.HeroDex += dex
//	return g.Update()
//}
//
//func (g *HeroTable) AddStar() error {
//	g.HeroStar += 1
//	return g.Update()
//}
//
//func (g *HeroTable) AddClass() error {
//	g.HeroClass += 1
//	return g.Update()
//}
