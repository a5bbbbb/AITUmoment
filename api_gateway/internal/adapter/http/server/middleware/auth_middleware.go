package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Middleware struct {
	jwtSecret string
}

func NewMiddleware(jwtSecret string) *Middleware {
	return &Middleware{jwtSecret: jwtSecret}
}

func (m *Middleware) AuthMiddleware(c *gin.Context) {

	tokenString, err := c.Cookie("auth_token")
	if err != nil {
		c.HTML(http.StatusUnauthorized, "index.html", gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenString, m.verifySigningMethod)
	if err != nil || !token.Valid {
		c.HTML(http.StatusUnauthorized, "index.html", gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Set("claims", claims)
	} else {
		c.HTML(http.StatusUnauthorized, "index.html", gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	c.Next()
}

func (m *Middleware) verifySigningMethod(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, http.ErrAbortHandler
	}
	return []byte(m.jwtSecret), nil
}
