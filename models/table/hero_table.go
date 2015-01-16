package table

import (
	"github.com/fhbzyc/c_game/libs/db"
)

func init() {
	db.DataBase.Sync(new(HeroTable))
}

var (
	HeroExpToLevel []int = []int{25, 55, 95, 155, 235}
)

type HeroTable struct {
	HeroId int  `xorm:"pk"`
	RoleId int  `xorm:"pk"`
	Exp    int  `xorm:"'hero_exp' NOT NULL DEFAULT 0"`
	Str    int  `xorm:"'hero_str' NOT NULL DEFAULT 0"`
	Int    int  `xorm:"'hero_int' NOT NULL DEFAULT 0"`
	Dex    int  `xorm:"'hero_dex' NOT NULL DEFAULT 0"`
	Star   int  `xorm:"'hero_star' NOT NULL DEFAULT 0 TINYINT(1)"`
	Class  int  `xorm:"'hero_class' NOT NULL DEFAULT 0 TINYINT(2)"`
	Skill1 int  `xorm:"'hero_skill_level1' NOT NULL DEFAULT 0"`
	Skill2 int  `xorm:"'hero_skill_level2' NOT NULL DEFAULT 0"`
	Skill3 int  `xorm:"'hero_skill_level3' NOT NULL DEFAULT 0"`
	Skill4 int  `xorm:"'hero_skill_level4' NOT NULL DEFAULT 0"`
	Equip1 bool `xorm:"'hero_equip1' NOT NULL DEFAULT 0 TINYINT(1)"`
	Equip2 bool `xorm:"'hero_equip2' NOT NULL DEFAULT 0 TINYINT(1)"`
	Equip3 bool `xorm:"'hero_equip3' NOT NULL DEFAULT 0 TINYINT(1)"`
	Equip4 bool `xorm:"'hero_equip4' NOT NULL DEFAULT 0 TINYINT(1)"`
	Equip5 bool `xorm:"'hero_equip5' NOT NULL DEFAULT 0 TINYINT(1)"`
	Equip6 bool `xorm:"'hero_equip6' NOT NULL DEFAULT 0 TINYINT(1)"`
}

func (h HeroTable) TableName() string {
	return "role_hero"
}

/*************************************************************************/

func (this *HeroTable) Level() int {
	for index := range HeroExpToLevel {
		if this.Exp <= HeroExpToLevel[index] {
			return index + 1
		}
	}
	return len(HeroExpToLevel)
}

func (this *HeroTable) GetPoint() int {
	return 3*this.Level() - this.Str - this.Int - this.Dex
}

func (this *HeroTable) GetSkillPoint() int {
	return 2 + 4 + 2*(this.Level()-1) - this.Skill1 - this.Skill2 - this.Skill3 - this.Skill4
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
