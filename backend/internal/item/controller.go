package item
import (
	"net/http"
	"github.com/gin-gonic/gin"
)
type Controller struct {
}

func NewController() Controller {
	return Controller{}
}

func (controller Controller) CreateItem(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success",
	})
}