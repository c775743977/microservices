package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-session/gin-session"
)

func main() {
	app := gin.Default()

	app.Use(ginsession.New())

	app.GET("/", func(ctx *gin.Context) {
		store := ginsession.FromContext(ctx)
		store.Set("foo", "bar")
		err := store.Save()
		if err != nil {
			ctx.AbortWithError(500, err)
			return
		}

		ctx.Redirect(302, "/foo")
	})

	app.GET("/foo", func(ctx *gin.Context) {
		store := ginsession.FromContext(ctx)
		foo, ok := store.Get("foo")
		if !ok {
			ctx.AbortWithStatus(404)
			return
		}
		ctx.String(http.StatusOK, "foo:%s", foo)
	})

	app.Run(":8080")
}