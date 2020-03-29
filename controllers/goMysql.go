package controllers

import (
	"github.com/astaxie/beego"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

type MysqlController struct {
	beego.Controller
}
//操作数据库的几个步骤
//数据库是不能直接用代码创建的，必须在终端创建，在代码中创建表
//1打开数据库
//2操作数据库
//3关闭数据库

func (this*MysqlController)ShowMysql(){
	conn, err := sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/class1?charset=utf8")
	if err != nil {
		beego.Info(err)
		return
	}
	//关闭数据库
    defer conn.Close()
	//操作数据库
	//创建数据库表
	//_,err =conn.Exec("CREATE TABLE userinfo(id int, name VARCHAR(11))")
	//if err != nil {
	//	beego.Info("Create Table is fialed",  err)
	//	return
	//}

	//往数据库表中插入数据
	_,err = conn.Exec("INSERT userinfo(id,name) VALUE (?,?)",1,"Tom")

	if err!= nil{
		beego.Info("插入数据错误",err)
		return
	}
	rows,err := conn.Query("select id from userinfo;")
	var id int
	for rows.Next() {
		rows.Scan(&id)
		beego.Info(id)
	}
	this.Ctx.WriteString("插入数据表成功")

}