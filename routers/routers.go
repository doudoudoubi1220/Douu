package routers

import (
	"github.com/gin-gonic/gin"
	"webproject01/controller"
)

func SetupRouters() *gin.Engine {
	r := gin.Default()
	//告诉gin框架模板文件引用的static文件去哪找
	r.Static("/static", "static")
	//告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)

	//v1
	v1Group := r.Group("v1")
	{
		v1Group.POST("/todo", controller.CreateATodo)
		v1Group.GET("/todo", controller.GetTodoList)
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return r
}
