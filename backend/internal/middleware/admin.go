package middleware

import (
	"net/http"

	"github.com/akekapong78/workflow/internal/constant"
	"github.com/gin-gonic/gin"
)

func CheckAdminRole(ctx *gin.Context) {
	// Get role from middleware
	if ctx.MustGet("role") != string(constant.UserRoleAdmin) {
		ctx.JSON(http.StatusForbidden, gin.H{
			"message": "only admin can update item status",
		})
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}

	ctx.Next()
}