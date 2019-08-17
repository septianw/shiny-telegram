package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingHandler(r *gin.Engine) {
	c.String(http.StatusOK, "pong")
}
