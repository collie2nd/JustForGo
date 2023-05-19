package main

import (
	_ "gin-swagger-demo/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type User struct {
}

func NewUser() User {
	return User{}
}

type StandardResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
}

// @Summary	禅道登录接口
// @Produce	json
// @Param		account		query		string				false	"用户名"
// @Param		password	query		string				false	"密码"
// @Success	200			{object}	StandardResponse	"成功"
// @Failure	400			{object}	StandardResponse	"请求错误"
// @Failure	500			{object}	StandardResponse	"内部错误"
// @Router		/api/v1/user/login [post]
func (t User) Login(c *gin.Context) {

}

// @Summary	禅道登出接口
// @Produce	json
// @Param		zentaosid	query		string				true	"session_id"
// @Success	200			{object}	StandardResponse	"成功"
// @Failure	400			{object}	StandardResponse	"请求错误"
// @Failure	500			{object}	StandardResponse	"内部错误"
// @Router		/api/v1/user/logout [post]
func (t User) Logout(c *gin.Context) {

}

// @title			helloworld
// @version		1.0
// @description	liuxing test helloworld
// @termsOfService	http://swagger.io/terms/
// @contact.name	lx
// @contact.url	http://www.swagger.io/support
// @contact.email	support@swagger.io
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			lx.helloworld.com
// @BasePath		/base/path
func main() {

	//注册swagger路由
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		//ginSwagger.DefaultModelsExpandDepth(-1),
		//ginSwagger.URL("/swagger/doc.json"),
	),
	)
	user := NewUser()
	apiv1 := r.Group("/api/v1")
	{
		u := apiv1.Group("user")
		{
			u.POST("/login", user.Login)
			u.POST("/logout", user.Logout)
		}
	}
	r.Run(":8080")
}
