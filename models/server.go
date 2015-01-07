package models

import (
	"github.com/fhbzyc/c_game/libs/db"
	"github.com/fhbzyc/c_game/models/table"
)

var Server serverModel

type serverModel struct {
}

func (this serverModel) FindAll() []table.ServerTable {
	serverList := make([]table.ServerTable, 0)
	db.DataBase.Where("server_enable = TRUE").Find(&serverList)
	return serverList
}
