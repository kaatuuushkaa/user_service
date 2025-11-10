package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
	appjwt "user_service/internal/jwt"
)

func JWTMiddleware(j appjwt.IJWT) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing Authorization header"})
			return
		}

		accessToken := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := j.ParseJWT(accessToken)
		if err != nil {
			if err == appjwt.ErrTokenExpired {
				refreshCookie, err := c.Cookie("REFRESH_TOKEN")
				if err != nil {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "refresh token not found"})
					return
				}

				newAccess, exp, err := j.RefreshAccessToken(refreshCookie)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "refresh token invalid"})
					return
				}

				c.Header("X-New-Access-Token", newAccess)
				c.SetCookie("ACCESS_TOKEN", newAccess, int(exp.Sub(time.Now()).Seconds()), "/", "", false, true)

				claims, err = j.ParseJWT(newAccess)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "failed to parse new access token"})
					return
				}
			}
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		c.Set("user_id", claims.ID)

		c.Next()
	}
}
