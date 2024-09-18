package item

import (
	"net/http"

	"github.com/akekapong78/workflow/internal/model"
	"github.com/akekapong78/workflow/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	Service  Service
	database *gorm.DB
}

func NewController(db *gorm.DB) Controller {
	return Controller{
		Service:  NewService(db),
		database: db,
	}
}

func (c Controller) CreateItem(ctx *gin.Context) {
	// Get user id
	userId, err := utils.GetUserIdByUsername(ctx.MustGet("username").(string), c.database)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
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

	item, err := c.Service.CreateItem(req, userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, item)
}

func (c Controller) GetItem(ctx *gin.Context) {
	// Get user id
	userId, err := utils.GetUserIdByUsername(ctx.MustGet("username").(string), c.database)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Path param
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	item, err := c.Service.GetItem(id, userId, ctx.MustGet("role").(string))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, item)
}

func (c Controller) GetItems(ctx *gin.Context) {
	// Get user id
	userId, err := utils.GetUserIdByUsername(ctx.MustGet("username").(string), c.database)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	items, err := c.Service.GetItems(userId, ctx.MustGet("role").(string))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, items)
}

func (c Controller) UpdateItem(ctx *gin.Context) {
	// Get user id
	userId, err := utils.GetUserIdByUsername(ctx.MustGet("username").(string), c.database)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

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

	item, err := c.Service.UpdateItem(id, req, userId, ctx.MustGet("role").(string))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, item)
}

func (c Controller) UpdateItemStatus(ctx *gin.Context) {
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

	item, err := c.Service.UpdateItemStatus(id, req, ctx.MustGet("role").(string))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, item)
}
