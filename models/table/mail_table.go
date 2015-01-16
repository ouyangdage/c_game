package table

import (
	"github.com/fhbzyc/c_game/libs/db"
)

func init() {
	db.DataBase.Sync(new(MailTable))
}

type MailTable struct {
	MailId    int    `xorm:"pk autoincr"`
	Type      int    `xorm:"'mail_type' NOT NULL DEFAULT 0 TINYINT(2)"`
	Title     string `xorm:"'mail_title' NOT NULL DEFAULT '' VARCHAR(100)"`
	Content   string `xorm:"'mail_content' NOT NULL DEFAULT '' VARCHAR(3000)"`
	Attchment string `xorm:"'mail_attchment' NOT NULL DEFAULT '' VARCHAR(3000)"`
	FromName  string `xorm:"'mail_from_name' NOT NULL DEFAULT '' VARCHAR(100)"`
	Time      string `xorm:"'mail_time' NOT NULL DEFAULT CURRENT_TIMESTAMP TIMESTAMP"`
	Status    int    `xorm:"'mail_status' NOT NULL DEFAULT 0 TINYINT(1)"`
}

func (this MailTable) TableName() string {
	return "role_mail"
}
