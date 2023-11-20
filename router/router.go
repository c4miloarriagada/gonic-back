package router

import (
	"encoding/gob"
	"go-gin-api/auth"
	"go-gin-api/web/callback"
	"go-gin-api/web/login"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func New(auth *auth.Authenticator) *gin.Engine {
	router := gin.Default()
	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	router.Static("/public", "index.html")
	router.LoadHTMLGlob("template/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", nil)
	})
	router.GET("/login", login.Handler(auth))
	router.GET("/redirect", callback.Handler(auth))
	// router.GET("/user", controllers.Handler)
	// router.GET("/logout", controllers.Handler)

	return router
}
