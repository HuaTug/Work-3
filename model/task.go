package model

import (
	"4-ToDoList/cache"
	"github.com/jinzhu/gorm"
	"strconv"
)

type Task struct {
	gorm.Model
	User      User   `gorm:"ForeignKey:Uid"` //建立外键，使得这两个表建立联系
	Uid       uint   `gorm:"not null"`
	Title     string `gorm:"index;not null"`
	Status    int    `gorm:"default:0"`
	Content   string `gorm:"type:longtext"`
	StartTime int64
	EndTime   int64 `gorm:"default:0"`
}

func (task *Task) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.TaskViewKey(task.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

// 使用redis缓存完成对于备忘录的访问次数
func (Task *Task) Addview() {
	cache.RedisClient.Incr(cache.TaskViewKey(Task.ID))
	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(Task.ID)))
}
