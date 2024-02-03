package main

import (
	"github.com/kataras/iris/v12"
)

type PingResponse struct {
	Message string `json:"message"`
}

func main() {
	app := iris.New()

	app.Get("/ping", func(ctx iris.Context) {
		res := PingResponse{
			Message: "pong",
		}

		ctx.JSON(res)
	})

	// this is to get url params
	app.Get("/{name}", func(ctx iris.Context) {
		params := ctx.Params()
		name := params.Get("name")

		ctx.JSON(name)
	})

	// this is to get url query
	app.Get("/welcome", func(ctx iris.Context) {
		firstname := ctx.URLParamDefault("firstname", "Guest")
		lastname := ctx.URLParam("lastname")

		ctx.Writef("Hello %s %s", firstname, lastname)
	})

	app.Listen(":8888")
}
