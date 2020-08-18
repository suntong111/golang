package models

import (
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
)

func init() {
	orm.Debug = true

	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:root@/text?charset=utf8",30)

	orm.RegisterModel(new(User),new(Article),new(Upload))

	orm.RunSyncdb("default", false, true)
}

