package controllers

import (
	"fmt"
	"github.com/fhbzyc/c_game/models"
	"github.com/fhbzyc/c_game/models/table"
)

func (this *Controller) Goods() error {

	typ := table.StoreType(this.Request.Params[0].(float64))

	goods := models.Store.Goods(typ, this.Connect.RoleId)
	return this.returnSuccess(goods)
}

// 刷新商店
func (this *Controller) Refresh() error {

	typ := table.StoreType(this.Request.Params[0].(float64))

	goods := models.Store.Refresh8num(typ, this.Connect.RoleId)
	return this.returnSuccess(goods)
}

// 买道具
func (this *Controller) Buy() error {

	typ := table.StoreType(this.Request.Params[0].(float64))
	index := int(this.Request.Params[1].(float64))

	if index < 1 || index > 8 {
		return this.returnError(lineNum(), ERROR_PARAM_ERROR)
	}

	goodsList := models.Store.Goods(typ, this.Connect.RoleId)

	goods := goodsList[index-1]

	if goods.Num <= 0 {
		return this.returnError(lineNum(), fmt.Errorf("道具已售馨"))
	}

	role := this.getRole()
	if goods.Type == table.GOODS_MONEY_COIN {
		if role.Coin < goods.Price {
			return this.returnError(lineNum(), ERROR_COIN_NOT_ENOUGH)
		} else {
			models.Role.SubCoin(role, goods.Price, table.FINANCE_BUY_COIN, "商店买<"+models.BaseItem.FindOne(goods.ItemId).Name)
		}
	} else if goods.Type == table.GOODS_MONEY_GOLD {
		if role.Coin < goods.Price {
			return this.returnError(lineNum(), ERROR_GOLD_NOT_ENOUGH)
		} else {
			models.Role.SubCoin(role, goods.Price, table.FINANCE_BUY_COIN, "商店买<"+models.BaseItem.FindOne(goods.ItemId).Name)
		}
	}

	item := models.Item.Add(role.RoleId, goods.ItemId, goods.Num)
	goodsList[index-1].Num = 0

	models.Store.Set(typ, role.RoleId, goodsList)

	return this.returnSuccess(item)
}
