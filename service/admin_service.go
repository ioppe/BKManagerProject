package service

import (
	"cmsproject/model"
	"github.com/go-xorm/xorm"
)

type AdminService interface {
	/*
	通过管理员用户名+密码 获取管理员实体，如果查询到，返回管理员实体，并返回true
	否则返回nil， false
	 */

	GetByAdminNameAndAdminPassword(username, password string) (model.Admin, bool)

	//返回管理员总数, 返回数值和布尔值
	GetAdminCount() (int64, error)
}

func NewAdminServie(db *xorm.Engine) AdminService {
	return &adminService{
		engine: db,
	}
}

type adminService struct {
	engine *xorm.Engine
}

func (as *adminService) GetAdminCount() (int64, error) {
	count, err := as.engine.Count(new(model.Admin))

	if err != nil {
		panic(err.Error())
		return 0, err
	}

	return count, nil
}

func (as *adminService) GetByAdminNameAndAdminPassword(username, password string) (model.Admin, bool) {
	var admin model.Admin

	as.engine.Where("admin_name = ? and pwd = ?", username, password).Find(&admin)

	return admin, admin.AdminId != 0
}
