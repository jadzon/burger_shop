package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jadzon/burger_shop/internal/app"
	"github.com/jadzon/burger_shop/internal/dto"
)

type UserHandler struct {
	app *app.Application
}

func NewUserHandler(app *app.Application) *UserHandler {
	return &UserHandler{
		app: app,
	}
}
func (h *UserHandler) Register(c *gin.Context) {
	var req dto.RegisterUserRequest
	err := c.ShouldBindJSON(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "failed to bind json"})
	}

}
