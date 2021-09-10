package datasource

import (
	"cmsproject/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func NewMysqlEngine() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:qwer1234@tcp(192.168.0.50:3306)/cmsProject?charset=utf8")
	if err != nil {
		panic(err)
	}

	err = engine.Sync2(new(
		model.Permission),
		new(model.City),
		new(model.Admin),
		new(model.AdminPermission),
		new(model.User),
		new(model.UserOrder))

	if err != nil {
		panic(err)
	}

	//设置显示sql语句
	engine.ShowSQL(true)
	engine.SetMaxOpenConns(10)

	return engine
}
