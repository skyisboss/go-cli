package router

import (
	"go-cli/internal/apps/web/handler"
	"go-cli/internal/apps/web/middleware"
	"go-cli/internal/ioc"

	"github.com/gin-gonic/gin"
)

func New(r *gin.Engine, ioc *ioc.Container) {
	r.Use(middleware.CorsMiddleware())

	handle := &handler.BaseHandler{
		Ioc: ioc,
	}
	v1 := r.Group("/api/v1/")
	{
		v1.GET("/demo", handle.Demo)
	}
}
