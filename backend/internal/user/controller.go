package user

import (
	"fmt"
	"net/http"
	"github.com/akekapong78/workflow/internal/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	Service Service
	domain  string
}

func NewController(db *gorm.DB, secret string, domain string) Controller {
	return Controller{
		Service: NewService(db, secret),
		domain:  domain,
	}
}

func (c Controller) GetUsers(ctx *gin.Context) {
	users, err := c.Service.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (c Controller) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	user, err := c.Service.GetUser(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c Controller) Register(ctx *gin.Context) {
	req := model.RequestUser{}

	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	user, err := c.Service.CreateUser(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (c Controller) Login(ctx *gin.Context) {
	req := model.RequestLogin{}

	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	token, err := c.Service.Login(&req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}	

	fmt.Println("generated token: ", token)

	// Retrieve cookie settings from environment variables or default values
	cookieDomain := c.domain
	if cookieDomain == "" {
		cookieDomain = "localhost"
	}
	
	ctx.SetCookie(
		"token",
		fmt.Sprintf("Bearer %v", token), 
		15 * 60, // 15 minutes 
		"/",
		cookieDomain,
		false,
		true,
	)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "login success",
	})
}