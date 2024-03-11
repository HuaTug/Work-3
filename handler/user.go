package handler

import (
	"4-ToDoList/service"
	"github.com/gin-gonic/gin"
)

// @Summary 用户注册
// @Produce json
// @Accept json
// @Param data body service.UserService true "用户名, 密码"
// @Success 200 {object} serializer.ResponseUser "{"status":200,"data":{},"msg":"ok"}"
// @Failure 500  {object} serializer.ResponseUser "{"status":500,"data":{},"Msg":{},"Error":"error"}"
// @Router /http://127.0.0.1:8000/api/user/login [post]
func UserRegister(c *gin.Context) {
	var user service.UserService
	if err := c.ShouldBind(&user); err == nil {
		res := user.Register()
		c.JSON(200, res)
	} else {
		c.JSON(400, "err")
	}
}

// @Summary 用户登录
// @Produce json
// @Accept json
// @Param     data    body     service.UserService    true      "user_name, password"
// @Success 200 {object} serializer.ResponseUser "{"success":true,"data":{},"msg":"登陆成功"}"
// @Failure 500 {object} serializer.ResponseUser "{"status":500,"data":{},"Msg":{},"Error":"error"}"
// @Router /user/login [post]
func UserLogin(c *gin.Context) {
	var userLoginService service.UserService
	if err := c.ShouldBind(&userLoginService); err == nil {
		res := userLoginService.Login()
		c.JSON(200, res)
	} else {
		c.JSON(400, "err")
	}
}
