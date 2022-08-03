package logging

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type RequestLog struct {
	StartTimeUnix  int64  `json:"start-time-unix"`
	DurationNano   int64  `json:"duration-nano"`
	Host           string `json:"host"`
	Method         string `json:"method"`
	Path           string `json:"path"`
	Query          string `json:"query"`
	Protocol       string `json:"protocol"`
	ResponseStatus int    `json:"response-status"`
	ResponseSize   int    `json:"response-size"`
	ServerName     string `json:"server-name"`
}

func LogJSON(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now().UnixNano()
		c.Next()
		stop := time.Now().UnixNano()

		log := RequestLog{
			StartTimeUnix:  start,
			DurationNano:   stop - start,
			Host:           c.Request.Host,
			Method:         c.Request.Method,
			Path:           c.Request.URL.EscapedPath(),
			Query:          c.Request.URL.RawQuery,
			Protocol:       c.Request.Proto,
			ResponseStatus: c.Writer.Status(),
			ResponseSize:   c.Writer.Size(),
			ServerName:     name,
		}

		raw, _ := json.Marshal(log)
		fmt.Println(string(raw))
	}
}
