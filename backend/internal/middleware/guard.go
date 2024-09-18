package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/akekapong78/workflow/internal/auth"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Guard(secret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// By header
		// authHeader := ctx.GetHeader("Authorization")

		// By cookie
		authCookie, err := ctx.Cookie("token")
		if err != nil {
			log.Println("Token missing in cookie")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
			return
		}
		// Remove prefix "Bearer " from auth token
		tokenString := strings.TrimPrefix(authCookie, "Bearer ")

		// Verify token
		token, err := auth.VerifyToken(tokenString, secret)
		if err != nil {
			log.Printf("Token verification failed: %v\\n", err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Extract claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Assuming 'aud' contains an array with username and role
			audience, ok := claims["aud"].([]interface{})
			if ok && len(audience) >= 2 {
				username := audience[0].(string)
				role := audience[1].(string)

				// Set the username and role in the context
				ctx.Set("username", username)
				ctx.Set("role", role)

				log.Printf("Token verified successfully. Username: %s, Role: %s\n", username, role)
			} else {
				log.Println("Invalid 'aud' claim format")
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		} else {
			log.Println("Invalid token claims")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Next()
	}
}
