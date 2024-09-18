package item

import (
	"fmt"
	"net/http"

	"github.com/akekapong78/workflow/internal/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	Service Service
}

func NewController(db *gorm.DB) Controller {
	return Controller{
		Service: NewService(db),
	}
}

func (c Controller) CreateItem(ctx *gin.Context) {
	req := model.RequestItem{}

	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	item, err := c.Service.CreateItem(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, item)
}

func (c Controller) GetItem(ctx *gin.Context) {
	userId := 1
	// Path param
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	item, err := c.Service.GetItem(id, userId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, item)
}

func (c Controller) GetItems(ctx *gin.Context) {
	userId := 1
	
	username, _ := ctx.Get("username")
	fmt.Println("username: ", username)

	items, err := c.Service.GetItems(userId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, items)
}

func (c Controller) UpdateItem(ctx *gin.Context) {
	userId := 1

	// Path param
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}	

	req := model.RequestItem{}
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}	

	item, err := c.Service.UpdateItem(id, req, userId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, item)
}

func (c Controller) UpdateItemStatus(ctx *gin.Context) {
	userId := 1

	// Path param
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	req := model.RequestUpdateItemStatus{}
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}	

	item, err := c.Service.UpdateItemStatus(id, req, userId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}	

	ctx.JSON(http.StatusOK, item)
}