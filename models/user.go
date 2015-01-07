package models

import (
	"github.com/fhbzyc/c_game/libs/db"
	"github.com/fhbzyc/c_game/models/table"
)

func init() {

}

var User userModel

type userModel struct {
}

func (this userModel) Insert(u *table.UserTable) error {
	_, err := db.DataBase.Insert(u)
	return err
}

func (this userModel) FineOne(uid int) (*table.UserTable, error) {
	u := new(table.UserTable)
	_, err := db.DataBase.Id(uid).Get(u)
	return u, err
}

func (this userModel) FindOneByPlatformId(platformId int, platformUuid string) (*table.UserTable, error) {
	u := new(table.UserTable)
	_, err := db.DataBase.Where("platform_id = ? AND platform_uuid = ?", platformId, platformUuid).Get(u)
	return u, err
}
