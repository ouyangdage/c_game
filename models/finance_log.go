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
