package controllers

import (
	"fmt"
	"github.com/fhbzyc/c_game/models"
	"github.com/fhbzyc/c_game/models/table"
)

func (this *Controller) HeroList() error {

	roleId := this.Connect.RoleId

	return this.returnSuccess(models.Hero.FindAll(roleId))
}

func (this *Controller) SkillUp() error {

	roleId := this.Connect.RoleId

	return this.returnSuccess(models.Hero.FindAll(roleId))
}

func (this *Controller) AddPoint() error {

	heroId := int(this.Request.Params[0].(float64))
	str := int(this.Request.Params[1].(float64))
	Int := int(this.Request.Params[2].(float64))
	dex := int(this.Request.Params[3].(float64))

	if str < 0 || Int < 0 || dex < 0 {
		return this.returnError(lineNum(), ERROR_PARAM_ERROR)
	}

	roleId := this.Connect.RoleId

	hero := getHero(roleId, heroId)

	if str+Int+dex > hero.GetPoint() {
		return this.returnError(lineNum(), fmt.Errorf("属性点超出上限"))
	}

	hero.Str += str
	hero.Int += Int
	hero.Dex += dex

	models.Hero.Update(hero)

	return this.returnSuccess(hero)
}

// sheng xing
func (this *Controller) StarUp() error {

	heroId := int(this.Request.Params[0].(float64))

	roleId := this.Connect.RoleId

	hero := getHero(roleId, heroId)

	if hero.Class >= 5 {
		return this.returnError(lineNum(), fmt.Errorf("已升级到5星"))
	}

	coin := 0
	num := 0

	if hero.Star <= 1 {
		num = 20
		coin = 35000
	} else if hero.Star == 2 {
		num = 50
		coin = 120000
	} else if hero.Star == 3 {
		num = 100
		coin = 300000
	} else {
		num = 150
		coin = 800000
	}

	role := this.getRole()
	if role.Coin < coin {
		return this.returnError(lineNum(), ERROR_COIN_NOT_ENOUGH)
	}

	item := getItem(roleId, heroId-6000)
	if item.Num < num {
		return this.returnError(lineNum(), fmt.Errorf("灵魂石数量不足"))
	}

	models.Role.SubCoin(role, coin, table.FINANCE_HERO_STAR_UP, fmt.Sprintf("<%s>升<%d>星", models.BaseHero.FindOne(heroId).Name, hero.Star))

	models.Item.Sub(item, num)

	hero.Star++
	models.Hero.Update(hero)

	return this.returnSuccess(hero)
}

// sheng jie
func (this *Controller) ClassUp() error {

	heroId := int(this.Request.Params[0].(float64))

	roleId := this.Connect.RoleId

	hero := getHero(roleId, heroId)

	hero.Class++

	models.BaseHeroEquip.FindOne(heroId, hero.Class)

	if !hero.Equip1 || !hero.Equip2 || !hero.Equip3 || !hero.Equip4 || !hero.Equip5 || !hero.Equip6 {
		return this.returnError(lineNum(), fmt.Errorf("装备不全不能升阶"))
	}

	models.Hero.Update(hero)

	return this.returnSuccess(hero)
}

// zhuang bei
func (this *Controller) Equipment() error {

	heroId := int(this.Request.Params[0].(float64))
	index := int(this.Request.Params[0].(float64))

	if index < 0 || index > 6 {
		return this.returnError(lineNum(), ERROR_PARAM_ERROR)
	}

	hero := this.getHero(heroId)
	heroEquip := models.BaseHeroEquip.FindOne(heroId, hero.Class)
	equipId := 0

	switch index {
	case 1:
		if hero.Equip1 {
			return this.returnError(lineNum(), fmt.Errorf("已配装备"))
		} else {
			hero.Equip1 = true
		}
		equipId = heroEquip.Equip1
	case 2:
		if hero.Equip2 {
			return this.returnError(lineNum(), fmt.Errorf("已配装备"))
		} else {
			hero.Equip2 = true
		}
		equipId = heroEquip.Equip2
	case 3:
		if hero.Equip3 {
			return this.returnError(lineNum(), fmt.Errorf("已配装备"))
		} else {
			hero.Equip3 = true
		}
		heroEquip := models.BaseHeroEquip.FindOne(heroId, hero.Class)
		equipId = heroEquip.Equip3
	case 4:
		if hero.Equip4 {
			return this.returnError(lineNum(), fmt.Errorf("已配装备"))
		} else {
			hero.Equip4 = true
		}
		equipId = heroEquip.Equip4
	case 5:
		if hero.Equip5 {
			return this.returnError(lineNum(), fmt.Errorf("已配装备"))
		} else {
			hero.Equip5 = true
		}
		equipId = heroEquip.Equip5
	case 6:
		if hero.Equip6 {
			return this.returnError(lineNum(), fmt.Errorf("已配装备"))
		} else {
			hero.Equip6 = true
		}
		equipId = heroEquip.Equip6
	}

	item := this.getItem(equipId)
	models.Item.Sub(item, 1)

	models.Hero.Update(hero)

	return this.returnSuccess(hero)
}

// chong zhi jie neng
func (this *Controller) ResetSkill() error {

	heroId := int(this.Request.Params[0].(float64))

	role := this.getRole()

	gold := 50

	if role.Gold < gold {
		return this.returnError(lineNum(), ERROR_GOLD_NOT_ENOUGH)
	} else {
		role.Gold -= gold
	}

	hero := this.getHero(heroId)
	hero.Skill1 = 0
	hero.Skill2 = 0
	hero.Skill3 = 0
	hero.Skill4 = 0

	models.Role.Update(role)
	models.Hero.Update(hero)

	return this.returnSuccess(hero)
}
