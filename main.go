package main

import (
	"4-ToDoList/conf"
	"4-ToDoList/routes"
)

func main() {
	conf.Init()
	r := routes.NewRouter()
	_ = r.Run(":10001")
}
