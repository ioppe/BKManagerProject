package model

import (
	"time"
)

type Admin struct {
	AdminId    int64     `xorm:"pk autoincr" json:"id"` //主键 自增
	AdminName  string    `xorm:"varchar(32)" json:"admin_name"`
	CreateTime time.Time `xorm:"DateTime" json:"create_time"`
	Status     int64     `xorm:"default 0" json:"status"`
	Avatar     string    `xorm:"varchar(255)" json:"avatar"`
	Pwd        string    `xorm:"varchar(255)" json:"pwd"`      //管理员密码
	CityName   string    `xorm:"varchar(12)" json:"city_name"` //管理员所在城市名称
	CityId     int64     `xorm:"index" json:"city_id"`
	City       *City     `xorm:"- <- ->"` //所对应的城市结构体（基础表结构体）
}

func (this *Admin) AdminToRespDesc() interface{} {
	return map[string]interface{}{
		"user_name": this.AdminName,
		"create_time": this.CreateTime,
		"status": this.Status,
		"avatar": this.Avatar,
		"pwd": this.Pwd,
		"city_name": this.CityName,
		"city_id": this.CityId,
		"admin": "管理员",
	}
}
