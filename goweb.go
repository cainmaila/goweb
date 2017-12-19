package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()
	app.StaticWeb("/", "./www")
	app.Run(iris.Addr(":80"))
}
