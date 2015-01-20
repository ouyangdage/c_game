package table

import (
	"github.com/fhbzyc/c_game/libs/db"
)

func init() {
	db.DataBase.Sync(new(HeroTable))
}

var (
	HeroExpToLevel []int = []int{25, 55, 95, 155, 235, 335, 465, 615, 795, 975, 1185, 1405, 1625, 1855, 2095, 2365, 2645, 2935, 3225, 3525, 3835, 4155, 4485, 4825, 5175, 5535, 5905,
		6285, 6675, 7075, 7605, 8285, 9115, 10095, 11195, 12595, 14295, 16295, 18595, 21195, 24195, 27495, 31195, 35295, 39795, 44795, 50195, 56095, 62495, 69495,
		76995, 85095, 93795, 103095, 113095, 124095, 135095, 147095, 160095, 174095, 189095, 204095, 220095, 237095, 255095, 275095, 296095, 318095, 341095, 365095,
		391095, 418095, 446095, 476095, 507095, 540095, 575095, 611095, 649095, 689095, 731095, 775095, 821095, 869095, 920095, 973095, 1028095, 1085095, 1144095, 1205095}
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
