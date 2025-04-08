package main

import (
	"miaoshaSystem/global"
	"miaoshaSystem/sql"
	"miaoshaSystem/web"
)

func main() {
	go global.StartKafkaConsumer()
	sql.Init()
	sql.ConnectMysql()
	global.CreateTable()
	web.Gin()

}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
