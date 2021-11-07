package http_server

import "github.com/gin-gonic/gin"

type Router interface {
	GinRouting(gin *gin.Engine)
}