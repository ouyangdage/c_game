package table

import (
	"github.com/fhbzyc/c_game/libs/db"
)

func init() {
	db.DataBase.Sync(new(RoleTable))
}

type RoleTable struct {
	RoleId           int    `xorm:"pk autoincr"`
	Name             string `xorm:"'role_name' index NOT NULL DEFAULT '' VARCHAR(10)"`
	Uid              int    `xorm:"NOT NULL DEFAULT 0 INT(11)"`
	AreaId           int    `xorm:"NOT NULL DEFAULT 0 INT(11)"`
	Coin             int    `xorm:"'role_coin' NOT NULL DEFAULT 0 INT(11)"`
	Gold             int    `xorm:"'role_gold' NOT NULL DEFAULT 0 INT(11)"`
	Exp              int    `xorm:"'role_exp' NOT NULL DEFAULT 0 INT(11)"`
	Progress         int    `xorm:"'role_progress' NOT NULL DEFAULT 0 INT(11)"`
	PhysicalStrength int    `xorm:"'role_physical_strength' NOT NULL DEFAULT 0 INT(11)"`
	PSTime           int    `xorm:"'role_ps_time' NOT NULL DEFAULT 0 INT(11)"`
}

func (this RoleTable) TableName() string {
	return "role"
}
