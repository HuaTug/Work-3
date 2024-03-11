package routes

import (
	"4-ToDoList/handler"
	"4-ToDoList/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	N := r.Group("api")
	{
		//用户操作
		N.POST("user/register", handler.UserRegister)
		N.POST("user/login", handler.UserLogin)
		authed := N.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.POST("task/create", handler.CreateTask)
			authed.PUT("task/update", handler.UpdateTask)
			authed.POST("task/search", handler.SearchTask)
			authed.POST("task/delete", handler.DeleteTask)
			authed.GET("task/list", handler.ListTask)
			authed.GET("task/show", handler.ShowTask)
		}
	}
	return r
}
