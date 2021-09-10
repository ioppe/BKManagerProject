package controller

import (
	"cmsproject/model"
	"cmsproject/service"
	"cmsproject/utils"
	"encoding/json"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

/*
管理员控制器
 */

type AdminController struct {
	//iris 框架自动为每个请求都绑定上下文对象
	Ctx iris.Context
	AdminService service.AdminService
	session *sessions.Session
}

const (
	ADMINTABLENAME = "admin"
	ADMIN = "admin"
)

type AdminUser struct {
	AdminName  string `json:"admin_name"` //管理员账户名
	Pwd        string `json:"pwd"`      //管理员密码
}

/**
管理员退出功能
请求类型： Get
请求url: admin/singout
 */

func (ac *AdminController) GetSingout() mvc.Result {
	//删除session 下次需要重新登录
	ac.session.Delete(ADMIN)
	return mvc.Response{
		Object: map[string]interface{}{
			"status": utils.RECODE_OK,
			"success": utils.Recode2Text(utils.RESPMSG_SIGNOUT),
		},
	}
}

/*
处理获取管理员总数的路由请求
请求类型：GET
请求url: admin/count
 */

func (ac *AdminController) GetCount() mvc.Result {
	count, err := ac.AdminService.GetAdminCount()
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status": utils.RECODE_OK,
				"message": utils.Recode2Text(utils.RESPMSG_ERRORADMINCOUNT),
			},
		}
	}

	return mvc.Response{
		Object: map[string]interface{}{
			"status": utils.RECODE_OK,
			"count": count,
		},
	}
}


/*
获取管理员信息接口
请求类型：GET
请求URL: admin/info
 */

func (ac *AdminController) GetInfo() mvc.Result {

	//从session中获取信息
	userByte := ac.session.Get(ADMIN)

	//session为空
	if userByte == nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status": utils.RECODE_UNLOGIN,
				"type": utils.EEROR_UNLOGIN,
				"message": utils.Recode2Text(utils.EEROR_UNLOGIN),
			},
		}
	}

	var admin model.Admin
	err := json.Unmarshal(userByte.([]byte), &admin)
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status": utils.RECODE_UNLOGIN,
				"type": utils.EEROR_UNLOGIN,
				"message": utils.Recode2Text(utils.EEROR_UNLOGIN),
			},
		}
	}

	return mvc.Response{
		Object: map[string]interface{}{
			"status": utils.RECODE_OK,
			"data": admin.AdminToRespDesc(),
		},
	}
}

/*
管理员登录功能
请求类型： POST
请求URL: admin/login
 */

func (ac *AdminController) PostLogin(ctx iris.Context) mvc.Result {

	iris.New().Logger().Info(" admin login")
	var adminUser AdminUser
	ac.Ctx.ReadJSON(&adminUser)

	if adminUser.AdminName == "" || adminUser.Pwd == "" {
		return mvc.Response{
			Object: map[string]interface{}{
				"status": "0",
				"success": "登录失败",
				"message": "用户名或密码为空，请重新填写后尝试登陆",
			},
		}
	}

	//根据用户名、密码到数据库中查询相应的管理信息
	admin, exist := ac.AdminService.GetByAdminNameAndAdminPassword(adminUser.AdminName, adminUser.Pwd)

	//管理员不存在
	if !exist {
		return mvc.Response{
			Object: map[string]interface{}{
				"status": "0",
				"success": "登陆失败",
				"message": "用户名或者密码错误，请重新登录",
			},
		}
	}

	//管理员存在，设置session
	userByte, _ := json.Marshal(admin)
	ac.session.Set(ADMIN, userByte)

	return mvc.Response{
		Object: map[string]interface{}{
			"status": "1",
			"success": "登陆成功",
			"message": "管理员登陆成功",
		},
	}
}
