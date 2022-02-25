package handler

import "github.com/gin-gonic/gin"

var Handler handlerInterface = &handler{}

type handlerInterface interface {
	CreateShortUrl(c *gin.Context)
	HandleShortUrlRedirect(c *gin.Context)
}

type handler struct {}
