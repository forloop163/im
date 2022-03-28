package main

import (
	"github.com/kataras/golog"
	"im-project/server/web"
	"os"
)



func init() {
	infoFile, err := os.OpenFile("logs/info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer infoFile.Close()

	golog.SetLevel("debug")
	golog.SetLevelOutput("info", infoFile)
}

func main() {
	// web 服务
	web.RunWebServer()
}
