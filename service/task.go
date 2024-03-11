package service

import (
	"4-ToDoList/model"
	"4-ToDoList/serializer"
	"time"
)

type ShowTaskService struct {
}
type CreateTaskService struct {
	Title   string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Content string `form:"content" json:"content" binding:"max=1000"`
	Status  int    `form:"status" json:"status"`
}

type UpdateTaskService struct {
	ID      uint   `form:"id" json:"id"`
	Title   string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Content string `form:"content" json:"content" binding:"max=1000"`
	Status  int    `form:"status" json:"status"`
}

type SearchTaskService struct {
	Info string `form:"info" json:"info"`
}

type ListTaskService struct {
	PageNum  int `json:"page_num" form:"page_num"`
	PageSize int `json:"page_size" form:"page_size"`
}

type DeleteTaskService struct {
}

func (service *SearchTaskService) Search(uId uint) serializer.Response {
	var tasks []model.Task
	model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uId).
		Where("title LIKE ? OR content LIKE ?",
			"%"+service.Info+"%", "%"+service.Info+"%").
		Find(&tasks)
	return serializer.Response{
		Status: 200,
		Msg:    " ",
		Data:   serializer.BuildTasks(tasks),
	}
}

func (service *CreateTaskService) Create(id uint) serializer.Response {
	var user model.User
	model.DB.First(&user, id)
	code := 200
	task := model.Task{
		User:      user,
		Uid:       user.ID,
		Title:     service.Title,
		Content:   service.Content,
		Status:    service.Status,
		StartTime: time.Now().Unix(),
		EndTime:   0,
	}
	err := model.DB.Create(&task).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "创建备忘录失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    "创建成功",
	}
}

func (service *UpdateTaskService) Update(id string) serializer.Response {
	var task model.Task
	model.DB.Model(model.Task{}).Where("uid=?", id).First(&task)
	task.Content = service.Content
	task.Status = service.Status
	task.Title = service.Title
	model.DB.Save(&task)
	return serializer.Response{
		Status: 200,
		Data:   serializer.BuildTask(task),
		Msg:    "更新完成",
	}
}

func (service *DeleteTaskService) Delete(id string) serializer.Response {
	var task model.Task
	err := model.DB.Delete(&task).Error
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "删除失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "删除成功",
	}
}

func (service *ListTaskService) List(id uint) serializer.Response {
	var tasks []model.Task
	var total int64
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	model.DB.Model(model.Task{}).Preload("User").Where("uid=?", id).Count(&total).
		Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).
		Find(&tasks)
	return serializer.BuildListResponse(tasks, uint(total))
}

func (service *ShowTaskService) Show(id uint) serializer.Response {
	var task model.Task
	code := 200
	err := model.DB.Where("uid=?", id).First(&task).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "查询失败",
		}
	}
	task.Addview()
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTask(task),
	}
}
