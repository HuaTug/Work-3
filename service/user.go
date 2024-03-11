package service

import (
	"4-ToDoList/model"
	"4-ToDoList/serializer"
	"4-ToDoList/utils"
	"github.com/jinzhu/gorm"
)

type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16"`
}

func (service *UserService) Register() *serializer.Response {
	var user model.User
	var count int64
	model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).First(&user).Count(&count)
	if count == 1 {
		return &serializer.Response{
			Status: 400,
			Msg:    "已经有人了，无需再注册",
		}
	}
	user.UserName = service.UserName
	if err := user.SetPassword(service.Password); err != nil {
		return &serializer.Response{
			Status: 400,
			Msg:    err.Error(),
		}
	}
	if err := model.DB.Create(&user).Error; err != nil {
		return &serializer.Response{
			Status: 500,
			Msg:    "数据库操作错误",
		}
	}
	return &serializer.Response{
		Status: 200,
		Msg:    "用户注册成功",
	}
}

func (service *UserService) Login() serializer.Response {
	var user model.User
	if err := model.DB.Where("user_name=?", service.UserName).First(&user).Error; err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return serializer.Response{
				Status: 400,
				Msg:    "用户不存在，请先注册",
			}

		}
		return serializer.Response{
			Status: 500,
			Msg:    "数据库错误",
		}
	}
	if !user.CheckPassword(service.Password) {
		return serializer.Response{
			Status: 400,
			Msg:    "密码错误",
		}
	}

	token, err := utils.GenerateToken(user.ID, service.UserName, 0)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "Token签发错误",
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    "登录成功",
	}
}
