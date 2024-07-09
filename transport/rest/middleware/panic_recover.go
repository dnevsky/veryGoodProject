package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
)

func PanicRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				buf := make([]byte, 8192)
				runtime.Stack(buf, false)
				panicHandler(fmt.Sprintf("%s\n%s", r, string(buf)))
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

func panicHandler(output string) {
	// ...
	fmt.Println(output)
}
