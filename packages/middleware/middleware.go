package middleware

import (
	jwttokengenrator "V3/packages/jwtTokenGenrator"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			log.Println("Error")
			c.JSON(401, gin.H{"error": "request does not contain an access token"})
			c.Abort()
			return
		}
		tokenSlice := strings.Split(authHeader, "Bearer ")
		if len(tokenSlice) < 2 {
			c.JSON(401, gin.H{"error": "invalid access token format"})
			c.Abort()
			return
		}
		tokenString := tokenSlice[1]
		claims, err := jwttokengenrator.ValidateToken(tokenString)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("uid", claims.Uid)
		c.Next()
	}
}
