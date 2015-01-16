package table

import (
	"github.com/fhbzyc/c_game/libs/db"
	"time"
)

func init() {
	db.DataBase.Sync(new(RoleTable))
}

var (
	ExpToLevel []int = []int{25, 50, 75, 105}
)

const (
	PSWaitTime = 6 * 60
)

type RoleTable struct {
	RoleId           int    `xorm:"pk autoincr"`
	Name             string `xorm:"'role_name' index NOT NULL DEFAULT '' VARCHAR(10)"`
	Uid              int    `xorm:"NOT NULL DEFAULT 0"`
	AreaId           int    `xorm:"NOT NULL DEFAULT 0"`
	Coin             int    `xorm:"'role_coin' NOT NULL DEFAULT 0"`
	Gold             int    `xorm:"'role_gold' NOT NULL DEFAULT 0"`
	Exp              int    `xorm:"'role_exp' NOT NULL DEFAULT 0"`
	Progress         int    `xorm:"'role_progress' NOT NULL DEFAULT 0"`
	PhysicalStrength int    `xorm:"'role_physical_strength' NOT NULL DEFAULT 0"`
	PSTime           int    `xorm:"'role_ps_time' NOT NULL DEFAULT 0"`
	Vip              int    `xorm:"'role_vip' NOT NULL DEFAULT 0"`
}

func (this RoleTable) TableName() string {
	return "role"
}

/**************************************************************************************/

func (this *RoleTable) GetLevel() int {
	for index, e := range ExpToLevel {
		if this.Exp < e {
			return index + 1
		}
	}
	return len(ExpToLevel)
}

func (this *RoleTable) MaxPhysicalStrength() int {
	return this.GetLevel() + 59
}

func (this *RoleTable) GetPhysicalStrength() int {

	n := (int(time.Now().Unix()) - this.PSTime) / PSWaitTime
	max := this.MaxPhysicalStrength()
	if n > max {
		n = max
	}
	return n + this.PhysicalStrength
}

func (this *RoleTable) SetPhysicalStrength(physicalStrength int) {

	ps := 0
	if physicalStrength > this.MaxPhysicalStrength() {
		ps = physicalStrength - this.MaxPhysicalStrength()
		physicalStrength = this.MaxPhysicalStrength()
	}

	now := int(time.Now().Unix())

	num := (now - this.PSTime) / PSWaitTime // num 个 时间恢复的体力

	time2 := now - num*PSWaitTime - this.PSTime // 多余的时间

	this.PSTime = now - (physicalStrength+1)*PSWaitTime + time2
	this.PhysicalStrength = ps
}
