package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Bootstrap() {
	log.Println("Ping module loaded.")
}

func Router(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}
