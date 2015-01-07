package models

import (
	"github.com/fhbzyc/c_game/libs/db"
	"github.com/fhbzyc/c_game/models/table"
)

func init() {
}

var GameArea gameAreaModel

type gameAreaModel struct {
}

func (this gameAreaModel) FindOne(areaId int) (table.GameAreaTable, error) {
	var gameArea table.GameAreaTable
	_, err := db.DataBase.Id(areaId).Get(&gameArea)
	return gameArea, err
}

func (this gameAreaModel) FindAll() ([]table.GameAreaTable, error) {
	var list []table.GameAreaTable
	return list, db.DataBase.Find(&list)
}

func (this gameAreaModel) CleanCache() error {
	return db.DataBase.ClearCache(new(table.GameAreaTable))
}
