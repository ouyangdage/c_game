package controllers

import (
	"errors"
	"github.com/fhbzyc/c_game/models"
	"github.com/fhbzyc/c_game/models/table"
	"github.com/fhbzyc/c_game/network"
	"github.com/fhbzyc/c_game/protocol"
)

type Controller struct {
	Connect *network.Connect
	Request *protocol.Request
}

var (
	ERROR_PARAM_ERROR     error = errors.New("参数错误")
	ERROR_COIN_NOT_ENOUGH error = errors.New("铜钱不足")
	ERROR_GOLD_NOT_ENOUGH error = errors.New("元宝不足")
	ERROR_HERO_NOT_HAVE   error = errors.New("没有这个英雄")
	ERROR_ITEM_NOT_HAVE   error = errors.New("没有这个道具")
)

func (this *Controller) getRole() *table.RoleTable {
	return getRole(this.Connect.RoleId)
}

func (this *Controller) getHero(heroId int) *table.HeroTable {
	return getHero(this.Connect.RoleId, heroId)
}

func (this *Controller) getItem(item int) *table.ItemTable {
	return getItem(this.Connect.RoleId, item)
}

func (this *Controller) itemList() []table.ItemTable {
	return itemList(this.Connect.RoleId)
}

func getRole(roleId int) *table.RoleTable {
	role, err := models.Role.FindOne(roleId)
	if err != nil {
		panic(err)
	}
	if role == nil {
		panic("角色id出错")
	}
	return role
}

func getHero(roleId, heroId int) *table.HeroTable {
	hero, err := models.Hero.FindOne(roleId, heroId)
	if err != nil {
		panic(err)
	}
	if hero == nil {
		panic(ERROR_HERO_NOT_HAVE)
	}
	return hero
}

func getItem(roleId, itemId int) *table.ItemTable {
	item, err := models.Item.FindOne(roleId, itemId)
	if err != nil {
		panic(err)
	}
	if item == nil {
		panic(ERROR_ITEM_NOT_HAVE)
	}
	return item
}

func itemList(roleId int) []table.ItemTable {
	items, err := models.Item.FindAll(roleId)
	if err != nil {
		panic(err)
	}
	return items
}
