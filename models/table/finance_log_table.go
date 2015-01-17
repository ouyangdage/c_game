package table

import (
	"github.com/fhbzyc/c_game/libs/db"
)

func init() {
	db.DataBase.Sync(new(FinanceLog))
}

type FinanceType int

const (
	_ FinanceType = iota
	FINANCE_ADMIN
	FINANCE_BUY_DIAMOND
	FINANCE_HERO_STAR_UP
	FINANCE_MAKE_ITEM
	FINANCE_BUY_COIN
	FINANCE_SELL_ITEM
	FINANCE_GENERAL_LEVELUP
	FINANCE_MAIL_GET
	FINANCE_DUPLICATE_USE
	FINANCE_DUPLICATE_GET
	FINANCE_SIGN_GET
	FINANCE_REVIVE
	FINANCE_GENERAL_LIST_ADD
)

type MoneyType int

const (
	TYPE_COIN MoneyType = 1
	TYPE_GOLD MoneyType = 2
)

type FinanceLog struct {
	RflId        int         `xorm:"pk autoincr"`
	RoleId       int         `xorm:"NOT NULL DEFAULT 0 INT(11)"`
	RflOldMoney  int         `xorm:"NOT NULL DEFAULT 0 INT(11)"`
	RflNewMoney  int         `xorm:"NOT NULL DEFAULT 0 INT(11)"`
	RflIsAdd     bool        `xorm:"NOT NULL DEFAULT 0 TINYINT(1)"`
	RflMoneyType MoneyType   `xorm:"NOT NULL DEFAULT 0 TINYINT(1)"`
	RflDesc      string      `xorm:"NOT NULL DEFAULT '' VARCHAR(255)"`
	RflStatic    FinanceType `xorm:"NOT NULL DEFAULT 0 TINYINT(2)"`
	RflTime      string      `xorm:"TIMESTAMP"`
}

func (this FinanceLog) TableName() string {
	return "role_finance_log"
}
