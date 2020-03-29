package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"TestDemo/models"
)

type OrmController struct {
	beego.Controller
}


func (this*OrmController)ShowOrm(){

	//查询表
	//获取一个连接对象 orm
	o:=orm.NewOrm()
	//获取查询对象
	user := &models.User{Id:1}

	err :=o.Read(user)
	if err != nil {
		beego.Info("查询失败",err)
		return
	}
	beego.Info(user.Name)
	this.Ctx.WriteString("查询成功")
}

func (this*OrmController)ShowInsert()  {


	//获取一个连接对象 orm
	o:=orm.NewOrm()
	//创建一个对象 id为主键不用赋值
	user := &models.User{}
	user.Name = "Jack"
	_, err :=o.Insert(user)
	if err != nil {
		beego.Info("插入表失败",err)
	}
	this.Ctx.WriteString("插入user表成功")


}

func (this*OrmController)ShowUpdate()  {
	////获取一个连接对象 orm
	o:=orm.NewOrm()
	////获取更新对象
	user:= &models.User{Id:1}
	//查询到对象赋值
	err :=o.Read(user)
	if err != nil {
		beego.Info("查询数据错误",err)
		return
	}

	// 给查询到的数据赋值
	user.Name = "Jarry"

	//更新操作
	_,err = o.Update(user)
	if err != nil {
		beego.Info("更新数据错误",err)
		return
	}
	this.Ctx.WriteString("更新数据成功")


}