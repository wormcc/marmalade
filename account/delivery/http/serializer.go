package http

import (
	"github.com/gin-gonic/gin"
	"github.com/wormcc/marmalade/models"
	"time"
)

type AccountSerializer struct {
	C *gin.Context
	*models.Account
}

type AccountResponse struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	CreateAt time.Time `json:"create_at"`
}

func (serializer *AccountSerializer) Response() AccountResponse {
	account := AccountResponse{
		ID:       serializer.ID,
		Name:     serializer.Name,
		Email:    serializer.Email,
		CreateAt: serializer.CreateAt,
	}
	return account
}
