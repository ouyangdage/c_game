package models

import (
	"github.com/fhbzyc/c_game/libs/db"
	"github.com/fhbzyc/c_game/models/table"
	"math/rand"
	"strconv"
	"time"
)

var Shop shopModel

type shopModel struct {
}

func (this shopModel) FindOne(roleId int) (*table.ShopTable, error) {
	shop := new(table.ShopTable)
	find, err := db.DataBase.Id(roleId).Get(shop)
	if err != nil {
		panic(err)
	}

	date, _ := strconv.Atoi(time.Now().Format("2006010215"))
	shop.Date = date
	shop.RoleId = roleId

	if !find {
		this.refresh(shop)
		db.DataBase.Insert(shop)
	} else if false {
		this.refresh(shop)
		this.Update(shop)
	}
	return shop, err
}

func (this shopModel) Update(shop *table.ShopTable) error {
	_, err := db.DataBase.Update(shop)
	return err
}

func (this shopModel) refresh(shop *table.ShopTable) {

	allGoods, _ := BaseStore.FindAll()

	find := make(map[int]int)

	var result []table.BaseStoreTable

	num := 8

	for {

		if len(find) >= num {
			break
		}

		temp := rand.Intn(len(allGoods))
		if _, ok := find[temp]; ok {
			continue
		} else {
			find[temp] = temp
		}

		result = append(result, allGoods[temp])
	}

	shop.Goods1 = result[0].GoodsId
	shop.Goods2 = result[1].GoodsId
	shop.Goods3 = result[2].GoodsId
	shop.Goods4 = result[3].GoodsId
	shop.Goods5 = result[4].GoodsId
	shop.Goods6 = result[5].GoodsId
	shop.Goods7 = result[6].GoodsId
	shop.Goods8 = result[7].GoodsId
	shop.IsBuy1 = false
	shop.IsBuy2 = false
	shop.IsBuy3 = false
	shop.IsBuy4 = false
	shop.IsBuy5 = false
	shop.IsBuy6 = false
	shop.IsBuy7 = false
	shop.IsBuy8 = false
}
