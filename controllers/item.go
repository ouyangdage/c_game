package controllers

import (
	"fmt"
	"github.com/fhbzyc/c_game/models"
	"github.com/fhbzyc/c_game/models/table"
)

func (this *Controller) ItemList() error {

	return this.returnSuccess(this.itemList())
}

// he cheng
func (this *Controller) MakeItem() error {

	itemId := int(this.Request.Params[0].(float64))

	baseItem := models.BaseItem.FindOne(itemId)

	role := this.getRole()

	if role.Coin < baseItem.MakeCoin {
		return this.returnError(lineNum(), ERROR_COIN_NOT_ENOUGH)
	}

	makeItems := models.BaseMakeItem.FindMaterial(baseItem.ItemId)

	var delItems []*table.ItemTable
	var numArray []int
	for _, makeItem := range makeItems {
		materia := this.getItem(makeItem.MaterialId)
		if materia.Num < makeItem.Num {
			return this.returnError(lineNum(), fmt.Errorf("材料不足"))
		} else {
			delItems = append(delItems, materia)
			numArray = append(numArray, makeItem.Num)
		}
	}

	models.Role.SubCoin(role, baseItem.MakeCoin, table.FINANCE_MAKE_ITEM, "合成<"+baseItem.Name)

	for index, materia := range delItems {
		models.Item.Sub(materia, numArray[index])
	}

	item := models.Item.Add(role.RoleId, itemId, 1)

	return this.returnSuccess(item)
}

// chu shou
func (this *Controller) SellItem() error {

	itemId := int(this.Request.Params[0].(float64))
	num := int(this.Request.Params[1].(float64))

	item := this.getItem(itemId)
	if item.Num < num {
		return this.returnError(lineNum(), fmt.Errorf("出售数量超过道具数量"))
	}

	models.Item.Sub(item, num)

	baseItem := models.BaseItem.FindOne(itemId)

	role := this.getRole()
	models.Role.AddCoin(role, baseItem.Coin, table.FINANCE_SELL_ITEM, "出售<"+baseItem.Name)

	return this.returnSuccess(role)
}
