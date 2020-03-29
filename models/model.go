package models

import (
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
)

type User struct {
	Id int
	Name string
}


func  init()  {
	//连接数据库
	orm.RegisterDataBase("default","mysql","root:123456@tcp(127.0.0.1:3306)/class1?charset=utf8")
	//注册表 有几个表写几个new参数
	orm.RegisterModel(new(User))
	//生成表(创建表) 参数：1数据库别名，2是否强制更新，3创建表过程是否可见
	orm.RunSyncdb("default",false,true)
}
