package middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

var jwtKey []byte
func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	secretKeyJwt := os.Getenv("JWT_KEY")
	jwtKey = []byte(secretKeyJwt)
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func JWTMiddleware() gin.HandlerFunc {
	return func (c *gin.Context){
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "tidak ada authorization header"})
			c.Abort()
			return
		}

		tokenString := string(authHeader)
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token)(interface{}, error){
			return jwtKey, nil
		})
		if err != nil || !token.Valid{
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token salah!"})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}

func GenerateJWT(username string) (string, error){
	expiredTime := time.Now().Add(1*time.Hour)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}