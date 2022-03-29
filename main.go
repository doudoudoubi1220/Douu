package main

import (
	"webproject01/dao"
	"webproject01/models"
	"webproject01/routers"
)

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//func initMySQL() (err error) {
//	dsn := "root:@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
//
//	DB, err = gorm.Open("mysql", dsn)
//	if err != nil {
//		return
//	}
//	//err=
//	return DB.DB().Ping()
//}

//type Todo struct {
//	ID     int64  `json:"id"`
//	Title  string `json:"title"`
//	Status bool   `json:"status"`
//}

func main() {
	//创建数据库
	//sql: CREATE DATABASE bubble
	//连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close() //程序退出关闭数据库连接

	//模型绑定
	dao.DB.AutoMigrate(&models.Todo{}) //todos

	r := routers.SetupRouters()

	//r := gin.Default()
	////告诉gin框架模板文件引用的static文件去哪找
	//r.Static("/static", "static")
	////告诉gin框架去哪里找模板文件
	//r.LoadHTMLGlob("templates/*")
	//r.GET("/", controller.IndexHandler)
	//
	////v1
	//v1Group := r.Group("v1")
	//{
	//	v1Group.POST("/todo", controller.CreateATodo)
	//	v1Group.POST("/todo", controller.GetTodoList)
	//	v1Group.POST("/todo/:id", controller.UpdateATodo)
	//	v1Group.POST("/todo/:id", controller.DeleteATodo)
	//}

	////待办事项
	////添加
	//v1Group.POST("/todo", func(c *gin.Context) {
	//	//从请求中拿出数据存到数据库
	//	//1.从请求中拿出数据
	//	var todo Todo
	//	c.BindJSON(&todo)
	//	//2.存入数据库
	//	err = DB.Create(&todo).Error
	//	if err != nil {
	//		fmt.Println("error", err)
	//	} else {
	//		c.JSON(http.StatusOK, gin.H{
	//			"code": 2000,
	//			"msg":  "success",
	//			"data": todo,
	//		})
	//	}
	//	//3.返回响应
	//
	//})
	////查看
	//v1Group.GET("/todo", func(c *gin.Context) {
	//	var todolist []Todo
	//	err = DB.Find(&todolist).Error
	//	if err != nil {
	//		c.JSON(http.StatusOK, gin.H{
	//			"error": err.Error(),
	//		})
	//	} else {
	//		c.JSON(http.StatusOK, todolist)
	//	}
	//})
	////修改某一个待办事项
	//v1Group.PUT("/todo/:id", func(c *gin.Context) {
	//	id, ok := c.Params.Get("id")
	//	if !ok {
	//		c.JSON(http.StatusOK, gin.H{"err": "无效的id"})
	//		return
	//	}
	//	var todo Todo
	//	err = DB.Where("id=?", id).First(&todo).Error
	//	if err != nil {
	//		c.JSON(http.StatusOK, gin.H{"err": err.Error()})
	//		return
	//	}
	//	c.BindJSON(&todo)
	//	if err = DB.Save(&todo).Error; err != nil {
	//		c.JSON(http.StatusOK, gin.H{
	//			"error": err.Error(),
	//		})
	//	} else {
	//		c.JSON(http.StatusOK, todo)
	//	}
	//
	//})
	////删除某一个待办事项
	//v1Group.DELETE("/todo/:id", func(c *gin.Context) {
	//	id, ok := c.Params.Get("id")
	//	if !ok {
	//		c.JSON(http.StatusOK, gin.H{"err": "无效的id"})
	//		return
	//	}
	//	if err = DB.Where("id=?", id).Delete(Todo{}).Error; err != nil {
	//		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	//	} else {
	//		c.JSON(http.StatusOK, gin.H{id: "delete"})
	//	}
	//})

	r.Run()
}
