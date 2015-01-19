package models

import (
	"encoding/json"
	"github.com/fhbzyc/c_game/libs/redis"
	"github.com/fhbzyc/c_game/models/table"
	"math/rand"
	"strconv"
	"time"
)

var (
	Store storeModel
)

type storeModel struct {
}

func (this storeModel) Key(typ table.StoreType, roleId int) string {
	return "SHOP_GOODS_" + strconv.Itoa(int(typ)) + "_" + strconv.Itoa(roleId)
}

func (this storeModel) Set(typ table.StoreType, roleId int, goods []table.BaseStoreTable) error {
	json, _ := json.Marshal(goods)
	return redis.Redis.Set(this.Key(typ, roleId), string(json))
}

func (this storeModel) Goods(typ table.StoreType, roleId int) []table.BaseStoreTable {

	var goods []table.BaseStoreTable

	str, err := redis.Redis.Get(this.Key(typ, roleId))
	if err != nil && err != redis.NotFind {
		panic(err)
	} else if err == redis.NotFind {
		goods = this.Refresh8num(typ, roleId)
	} else {
		json.Unmarshal([]byte(str), &goods)

		if time.Now().Unix() >= goods[0].Time {
			goods = this.Refresh8num(typ, roleId)
		}
	}

	return goods
}

func (this storeModel) refresh(typ table.StoreType) []table.BaseStoreTable {

	allGoods, _ := BaseStore.FindAll(typ)

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

		now := time.Now()
		todayBeginTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Unix()

		time1 := todayBeginTime + 9*3600
		time2 := todayBeginTime + 12*3600
		time3 := todayBeginTime + 18*3600
		time4 := todayBeginTime + 21*3600

		t := int64(0)
		if now.Unix() < time1 {
			t = time1
		} else if now.Unix() < time2 {
			t = time2
		} else if now.Unix() < time3 {
			t = time3
		} else if now.Unix() < time4 {
			t = time4
		} else {
			t = time1 + 86400
		}

		goods := allGoods[temp]
		goods.Time = t

		result = append(result, allGoods[temp])
	}

	return result
}

func (this storeModel) Refresh8num(typ table.StoreType, roleId int) []table.BaseStoreTable {

	goods := this.refresh(typ)

	this.Set(typ, roleId, goods)

	return goods
}
