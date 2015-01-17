package table

import (
	"github.com/fhbzyc/c_game/libs/db"
)

func init() {
	db.DataBase.Sync(new(MailTable))
}

type MailTable struct {
	MailId    int    `xorm:"pk autoincr"`
	RoleId    int    `xorm:"index NOT NULL DEFAULT 0"`
	Type      int    `xorm:"'mail_type' NOT NULL DEFAULT 0 TINYINT(1)"`
	Title     string `xorm:"'mail_title' NOT NULL DEFAULT ''"`
	Content   string `xorm:"'mail_content' NOT NULL DEFAULT '' VARCHAR(1000)"`
	Attchment string `xorm:"'mail_attchment' NOT NULL DEFAULT '' VARCHAR(1000)"`
	FromName  string `xorm:"'mail_from_name' NOT NULL DEFAULT ''"`
	Time      string `xorm:"'mail_time' NOT NULL DEFAULT CURRENT_TIMESTAMP TIMESTAMP"`
	Status    int    `xorm:"'mail_status' NOT NULL DEFAULT 0 TINYINT(1)"`
}

func (this MailTable) TableName() string {
	return "role_mail"
}
