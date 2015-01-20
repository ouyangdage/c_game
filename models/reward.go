package models

import (
	"github.com/fhbzyc/c_game/models/table"
)

type Reward struct {
	Items []*table.ItemTable
	Heros []*table.HeroTable
	Coin  int
	Gold  int
}

func (this *Reward) SetCoin(coin int) {
	this.Coin = coin
}

func (this *Reward) SetGold(gold int) {
	this.Gold = gold
}

func (this *Reward) AddItem(item *table.ItemTable) {
	this.Items = append(this.Items, item)
}

func (this *Reward) AddHero(hero *table.HeroTable) {
	this.Heros = append(this.Heros, hero)
}
