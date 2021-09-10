package main

import (
	"cmsproject/config"
	"cmsproject/controller"
	"cmsproject/datasource"
	"cmsproject/service"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"time"
)

//工程入口
func main()  {
	app := NewApp()

	//应用App设置
	Configuration(app)

	//路由设置
	MvcHandle(app)

	config := config.InitConfig()
	addr := ":"+config.Port
	app.Run(iris.Addr(addr))
}

func NewApp() *iris.Application {
	app := iris.New()

	app.Logger().SetLevel("debug")


	app.HandleDir("/static","./static")
	app.HandleDir("/manage/static/", "./static")
	app.HandleDir("/img", "./static/img")

	app.RegisterView(iris.HTML("./static", ".html").Reload(true))
	app.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})

	return app
}

/*
MVC架构处理
 */
func MvcHandle(app *iris.Application) {
	//启用session
	sessManager := sessions.New(sessions.Config{
		Cookie: "sessionCookie",
		Expires: 24 * time.Hour,
	})

	engine := datasource.NewMysqlEngine()

	//管理员模块功能
	adminService := service.NewAdminServie(engine)

	admin := mvc.New(app.Party("/admin"))
	admin.Register(
		adminService,
		sessManager.Start,
		)
	admin.Handle(new(controller.AdminController))

	//用户功能模块


	//商品功能模块
}

/*
项目设置
 */
func Configuration(app *iris.Application) {

	//配置字符编码
	app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset: "UTF-8",
	}))

	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"errmsg": iris.StatusNotFound,
			"msg": " not found",
			"data": iris.Map{},
		})
	})

	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"errmsg": iris.StatusInternalServerError,
			"msg": "internal Error",
			"data": iris.Map{},
		})
	})
}