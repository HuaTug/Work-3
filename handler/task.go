package handler

import (
	"4-ToDoList/service"
	"4-ToDoList/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Summary 创建任务
// @Produce json
// @Accept json
// @Header 200 {string} token "必备"
// @Param data body service.CreateTaskService true "title,content,status"
// @Success 200 {object} serializer.Response "{"status":200,"data":{},"msg":"ok"}"
// @Failure 500 {object} serializer.Response "{"status":200,"data":{},"msg":"ok"}"
// @Router /api/task/create [post]
func CreateTask(c *gin.Context) {
	var createTask service.CreateTaskService
	claim, _ := utils.ParseToken(c.GetHeader("token"))
	if err := c.ShouldBind(&createTask); err == nil {
		res := createTask.Create(claim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, "err")
	}
}

// @Summary 修改任务
// @Produce json
// @Accept json
// @Header 200 {string} token "必备"
// @Param data body service.UpdateTaskService true "id,title,content"
// @Success 200 {object} serializer.Response "{"status":200,"data":{},"msg":"ok"}"
// @Failure 500 {object} serializer.Response "{"status":200,"data":{},"msg":"ok"}"
// @Router /api/task/update [put]
func UpdateTask(c *gin.Context) {
	var update service.UpdateTaskService
	claim, _ := utils.ParseToken(c.GetHeader("token"))
	if err := c.ShouldBind(&update); err == nil {
		res := update.Update(strconv.Itoa(int(claim.Id)))
		c.JSON(200, res)
	} else {
		c.JSON(400, "err")
	}
}

// @Summary 查询任务
// @Produce json
// @Accept json
// @Header 200 {string} token "必备"
// @Param data body service.SearchTaskService true "id"
// @Success 200 {object} serializer.Response "{"status":200,"data":{},"msg":"ok"}"
// @Failure 500 {object} serializer.Response "{"status":200,"data":{},"msg":"ok"}"
// @Router /api/task/search [post]
func SearchTask(c *gin.Context) {
	var search service.SearchTaskService
	claim, _ := utils.ParseToken(c.GetHeader("token"))
	if err := c.ShouldBind(&search); err == nil {
		res := search.Search(claim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, "err")
	}
}

// @Summary 删除任务
// @Produce json
// @Accept json
// @Header 200 {string} token "必备"
// @Param data body service.DeleteTaskService true "id"
// @Success 200 {object} serializer.Response "{"status":200,"data":{},"msg":"ok"}"
// @Failure 500 {object} serializer.Response "{"status":200,"data":{},"msg":"ok"}"
// @Router /api/task/delete [delete]
func DeleteTask(c *gin.Context) {
	var delete service.DeleteTaskService
	claim, _ := utils.ParseToken(c.GetHeader("token"))
	if err := c.ShouldBind(&delete); err == nil {
		res := delete.Delete(strconv.Itoa(int(claim.Id)))
		c.JSON(200, res)
	} else {
		c.JSON(400, "err")
	}
}

// @Summary 展示所有任务
// @Produce json
// @Accept json
// @Header 200 {string} token "必备"
// @Param data body service.ListTaskService true "token"
// @Success 200 {object} serializer.Response "{"status":200,"data":{},"msg":"ok"}"
// @Failure 500 {object} serializer.Response "{"status":200,"data":{},"msg":"ok"}"
// @Router /api/task/list [get]
func ListTask(c *gin.Context) {
	var list service.ListTaskService
	claim, _ := utils.ParseToken(c.GetHeader("token"))
	if err := c.ShouldBind(&list); err == nil {
		res := list.List(claim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, "err")
	}
}

// @Summary 展示任务
// @Produce json
// @Accept json
// @Header 200 {string} token "必备"
// @Param data body service.ShowTaskService true "token"
// @Success 200 {object} serializer.Response "{"status":200,"data":{},"msg":"ok"}"
// @Failure 500 {object} serializer.Response "{"status":200,"data":{},"msg":"ok"}"
// @Router /task/show [get]
func ShowTask(c *gin.Context) {
	var show service.ShowTaskService
	Claim, _ := utils.ParseToken(c.GetHeader("token"))
	if err := c.ShouldBind(&show); err == nil {
		res := show.Show(Claim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, "err")
	}
}
