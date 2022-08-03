package main

import (
	"fmt"
	"go-credit-rate-limit-server/pkg/logging"
	"go-credit-rate-limit-server/pkg/rate"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RequestData struct {
	Account string `form:"account" json:"account" binding:"required"`
	Credits int    `form:"credits" json:"credits" binding:"required"`
}

func health(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

func acquire(l *rate.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request RequestData
		err := c.Bind(&request)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		acquired, err := l.Acquire(request.Account, request.Credits)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"credits": acquired,
		})
	}
}

func main() {
	limiter := &rate.Limiter{
		Accounts: map[string]rate.Rate{
			"example": rate.NewRate(5*time.Second, 15),
		},
	}

	r := gin.New()

	r.UseRawPath = true

	r.Use(gin.Recovery())

	r.GET("/health", health)

	v1 := r.Group("/v1")
	{
		v1.Use(logging.LogJSON("credit-rate-limit-server"))
		v1.GET("/acquire", acquire(limiter))
	}

	err := r.Run()
	if err != nil {
		fmt.Println("failed with error:", err.Error())
	}
}
