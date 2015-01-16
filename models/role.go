package models

import (
	"github.com/fhbzyc/c_game/libs/db"
	"github.com/fhbzyc/c_game/models/table"
	"time"
)

var Role roleModel = roleModel{
	LevelUpPS: []int{10, 10, 20, 20, 20, 20},
}

type roleModel struct {
	LevelUpPS []int
}

func (this roleModel) FindOneByArea(uid, areaId int) (*table.RoleTable, error) {
	Role := new(table.RoleTable)
	_, err := db.DataBase.Where("uid = ? AND area_id = ? ", uid, areaId).Get(Role)
	return Role, err
}

func (this roleModel) FindOne(roleId int) (*table.RoleTable, error) {
	Role := new(table.RoleTable)
	_, err := db.DataBase.Id(roleId).Get(Role)
	return Role, err
}

func (this roleModel) Insert(uid, areaId int) (*table.RoleTable, error) {
	Role := new(table.RoleTable)
	Role.Uid = uid
	Role.AreaId = areaId
	_, err := db.DataBase.Insert(Role)
	return Role, err
}

func (this roleModel) AddCoin(Role *table.RoleTable, FinanceType table.FinanceType, coin int, desc string) error {

	oldmoney := Role.Coin

	Role.Coin += coin
	err := this.Update(Role)
	if err == nil {
		finance := &table.FinanceLog{
			RoleId:       Role.RoleId,
			RflOldMoney:  oldmoney,
			RflNewMoney:  Role.Coin,
			RflIsAdd:     true,
			RflMoneyType: table.TYPE_COIN,
			RflDesc:      desc,
			RflStatic:    FinanceType,
			RflTime:      time.Now().Format("2006-01-02 15:04:05"),
		}
		InsertFinanceLog(finance)
	}
	return err
}

func (this roleModel) SubCoin(Role *table.RoleTable, FinanceType table.FinanceType, coin int, desc string) error {

	oldmoney := Role.Coin

	Role.Coin -= coin
	err := this.Update(Role)
	if err == nil {
		finance := &table.FinanceLog{
			RoleId:       Role.RoleId,
			RflOldMoney:  oldmoney,
			RflNewMoney:  Role.Coin,
			RflIsAdd:     false,
			RflMoneyType: table.TYPE_COIN,
			RflDesc:      desc,
			RflStatic:    FinanceType,
			RflTime:      time.Now().Format("2006-01-02 15:04:05"),
		}
		InsertFinanceLog(finance)
	}
	return err
}

func (this roleModel) AddGold(Role *table.RoleTable, FinanceType table.FinanceType, gold int, desc string) error {

	oldmoney := Role.Gold

	Role.Gold += gold
	err := this.Update(Role)
	if err == nil {
		finance := &table.FinanceLog{
			RoleId:       Role.RoleId,
			RflOldMoney:  oldmoney,
			RflNewMoney:  Role.Gold,
			RflIsAdd:     true,
			RflMoneyType: table.TYPE_GOLD,
			RflDesc:      desc,
			RflStatic:    FinanceType,
			RflTime:      time.Now().Format("2006-01-02 15:04:05"),
		}
		InsertFinanceLog(finance)
	}
	return err
}

func (this roleModel) SubGold(Role *table.RoleTable, FinanceType table.FinanceType, gold int, desc string) error {

	oldmoney := Role.Gold

	Role.Gold -= gold
	err := this.Update(Role)
	if err == nil {
		finance := &table.FinanceLog{
			RoleId:       Role.RoleId,
			RflOldMoney:  oldmoney,
			RflNewMoney:  Role.Gold,
			RflIsAdd:     false,
			RflMoneyType: table.TYPE_GOLD,
			RflDesc:      desc,
			RflStatic:    FinanceType,
			RflTime:      time.Now().Format("2006-01-02 15:04:05"),
		}
		InsertFinanceLog(finance)
	}
	return err
}

func (this roleModel) Count(name string) (int, error) {
	n, err := db.DataBase.Where("role_name = ?", name).Count(new(table.RoleTable))
	return int(n), err
}

func (this roleModel) AddExp(role *table.RoleTable, exp int) error {

	//	oldExp := role.Exp
	oldLevel := role.GetLevel()
	oldPS := role.GetPhysicalStrength()

	maxExp := table.ExpToLevel[len(table.ExpToLevel)-1]

	role.Exp += exp
	if role.Exp > maxExp {
		role.Exp = maxExp
	}

	newLevel := role.GetLevel()

	if newLevel > oldLevel {
		addPS := 0
		for i := 0; i > oldLevel; i-- {
			addPS += this.LevelUpPS[i-2]
		}

		role.SetPhysicalStrength(addPS + oldPS)

		this.LevelUp(role)
	}

	return nil
}

func (this roleModel) LevelUp(role *table.RoleTable) {
	go func() {
		// 检查升级主线任务

	}()
}

func (this roleModel) Update(role *table.RoleTable) error {
	_, err := db.DataBase.Update(role)
	return err
}

//
//func (this *RoleData) ActionValue() int {
//
//	now := time.Now()
//	n := (int(now.Unix() - this.ActionTime)) / Role.ActionWaitTime
//	if n > Role.MaxActionValue {
//		n = Role.MaxActionValue
//	}
//
//	return int(n) + this.OtherAction
//}
//
//func (this *RoleData) SetActionValue(n int) error {
//
//	RoleData := *this
//
//	if n > Role.MaxActionValue {
//		this.OtherAction = n - Role.MaxActionValue
//		n = Role.MaxActionValue
//	} else {
//		this.OtherAction = 0
//	}
//
//	nowUnix := time.Now().Unix()
//	remainder := int(nowUnix-this.ActionTime) % Role.ActionWaitTime
//	this.ActionTime = nowUnix - int64(remainder) - int64(Role.ActionWaitTime*n)
//
//	_, err := DataBase.Update(this)
//	if err != nil {
//		this = &RoleData
//	}
//	return err
//}
//
//func (this *RoleData) BuyActionValue(diamond int, n int) error {
//
//	RoleData := *this
//
//	if n > Role.MaxActionValue {
//		this.OtherAction = n - Role.MaxActionValue
//		n = Role.MaxActionValue
//	}
//
//	nowUnix := time.Now().Unix()
//	remainder := int(nowUnix-this.ActionTime) % Role.ActionWaitTime
//	this.ActionTime = nowUnix - int64(remainder) - int64(Role.ActionWaitTime*n)
//	this.Diamond -= diamond
//
//	_, err := DataBase.Update(this)
//	if err == nil {
//		//		InsertSubDiamondFinanceLog(this.Uid, FINANCE_BUY_ACTION, oldDiamond, this.Diamond, fmt.Sprintf("%d -> %d", oldAction, n))
//	} else {
//		this = &RoleData
//	}
//	return err
//}
//
//func (this *RoleData) ActionRecoverTime() int {
//
//	nowUnix := time.Now().Unix()
//	remainder := int(nowUnix-this.ActionTime) % Role.ActionWaitTime
//
//	return Role.ActionWaitTime - remainder
//}
//
//func (this *RoleData) SetGeneralBaseId(baseId int) error {
//
//	temp := this.GeneralBaseId
//	this.GeneralBaseId = baseId
//
//	_, err := DataBase.Update(this)
//	if err != nil {
//		this.GeneralBaseId = temp
//	}
//	return err
//}
//
//func (this *RoleData) Sign() error {
//
//	temp1 := this.SignDate
//	temp2 := this.SignNum
//	now := time.Now()
//
//	if this.SignDate == now.Format("20060102") {
//		return nil
//	}
//
//	if this.SignDate == now.AddDate(0, 0, -1).Format("20060102") {
//		this.SignNum++
//	} else {
//		this.SignNum = 1
//	}
//
//	this.SignDate = now.Format("20060102")
//
//	_, err := DataBase.Update(this)
//	if err != nil {
//		this.SignDate = temp1
//		this.SignNum = temp2
//	}
//
//	return err
//}
//
//func (this *RoleData) UpdateDate() error {
//
//	now := time.Now()
//	if this.BuyActionDate == now.Format("20060102") {
//		return nil
//	}
//
//	temp1 := this.BuyActionDate
//	temp2 := this.BuyActionNum
//
//	this.BuyActionDate = now.Format("20060102")
//	this.BuyActionNum = 0
//
//	_, err := DataBase.Update(this)
//	if err != nil {
//		this.BuyActionDate = temp1
//		this.BuyActionNum = temp2
//	}
//
//	return err
//}
//
//func (this *RoleData) Set() error {
//	_, err := DataBase.Update(this)
//	return err
//}
//
//func (this RoleModel) NewRole(uid int64) (*RoleData, error) {
//
//	RoleData := new(RoleData)
//
//	return RoleData, nil
//}
