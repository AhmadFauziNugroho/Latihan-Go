package middleware

import(
	"net/http"
	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	c.Next()
	
	if len(c.Errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": c.Errors[0].Error()})
	}
}