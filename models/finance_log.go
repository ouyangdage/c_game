package models

import (
	"github.com/fhbzyc/c_game/libs/db"
	"github.com/fhbzyc/c_game/libs/log"
	"github.com/fhbzyc/c_game/models/table"
)

var (
	financeStrMap map[table.FinanceType]string = make(map[table.FinanceType]string)
	financeChan   chan *table.FinanceLog       = make(chan *table.FinanceLog, 1000)
)

func init() {
	go readFinanceChan()

	financeStrMap[table.FINANCE_ADMIN] = "内部添加"
	financeStrMap[table.FINANCE_BUY_DIAMOND] = "充值"
	financeStrMap[table.FINANCE_BUY_ACTION] = "买体力"
	financeStrMap[table.FINANCE_BUY_GENERAL] = "买英雄"
	financeStrMap[table.FINANCE_BUY_COIN] = "兑换金币"
	financeStrMap[table.FINANCE_ITEM_LEVELUP] = "道具升级"
	financeStrMap[table.FINANCE_GENERAL_LEVELUP] = "英雄升级"
	financeStrMap[table.FINANCE_MAIL_GET] = "邮件领取"
	financeStrMap[table.FINANCE_DUPLICATE_USE] = "临时道具"
	financeStrMap[table.FINANCE_DUPLICATE_GET] = "副本获得"
	financeStrMap[table.FINANCE_SIGN_GET] = "签到"
	financeStrMap[table.FINANCE_REVIVE] = "复活"
	financeStrMap[table.FINANCE_GENERAL_LIST_ADD] = "添加英雄"
}

func readFinanceChan() {

	defer func() {
		if err := recover(); err != nil {
			log.Logger.Error("financeChan panic : %v", err)
			readFinanceChan()
		}
	}()

	for finance := range financeChan {
		db.DataBase.Insert(finance)
		//sql := "INSERT INTO g_role_finance_log SET roles_unique = ? , rfl_old_money = ? , rfl_new_money = ? , rfl_type = ? , rfl_mtype = ? , rfl_desc = ? , rfl_time = UNIX_TIMESTAMP() , rfl_static_type = ? "
		//DB().Exec(sql, finance.roleId, finance.oldMoney, finance.newMoney, isAdd, finance.moneyType, finance.desc, finance.FinanceType)
	}
}

func InsertFinanceLog(financeLog *table.FinanceLog) {
	financeChan <- financeLog
}
