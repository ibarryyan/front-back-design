package main

import (
	"back_go/config"
	"back_go/model"
	"back_go/web"
)

func main() {
	config.InitDB()
	model.DefaultStudentDb = model.NewStudentDb()
	web.RunHttp()
}
