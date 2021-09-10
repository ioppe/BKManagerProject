package model

import "time"

type User struct {
	Id           int64     `xorm:"pk autoincr" json:"id"`        //主键 用户ID
	UserName     string    `xorm:"varchar(12)" json:"username"`  //用户名称
	RegisterTime time.Time `json:"register_time"`                //用户注册时间
	Mobile       string    `xorm:"varchar(11)" json:"mobile"`    //用户的移动手机号
	IsActive     int64     `json:"is_active"`                    //用户是否激活
	Balance      int64     `json:"balance"`                      //用户的账户余额（简单起见，使用int类型）
	Avatar       string    `xorm:"varchar(255)" json:"avatar"`   //用户的头像
	Pwd          string    `json:"password"`                     //用户的账户密码
	DelFlag      int64     `json:"del_flag"`                     //是否被删除的标志字段 软删除
	CityName     string    `xorm:"varchar(24)" json:"city_name"` //用户所在城市的名称
	City         *City     `xorm:"- <- ->"`
}


