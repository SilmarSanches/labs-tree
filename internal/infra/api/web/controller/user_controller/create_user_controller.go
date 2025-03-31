package user_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"fullcycle-auction_go/internal/entity/user_entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *user_entity.User) error
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var user user_entity.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	user.Id = uuid.New().String()

	err := uc.userUseCase.CreateUser(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar usuário"})
		return
	}

	c.JSON(http.StatusCreated, user)
}
