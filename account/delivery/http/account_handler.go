package http

import (
	"github.com/gin-gonic/gin"
	"github.com/wormcc/marmalade/account"
	"strconv"
)

type AccountHttpHandler struct {
	AccountUseCase account.UseCase
}

func NewAccountHttpHandler(r *gin.Engine, us account.UseCase) {
	handler := &AccountHttpHandler{
		AccountUseCase: us,
	}
	accountGroup := r.Group("/users")
	{
		accountGroup.GET("/:id", handler.GetAccount)
		accountGroup.POST("/", handler.CreateAccount)
	}
}

func (h *AccountHttpHandler) CreateAccount(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Created!"})
}

func (h *AccountHttpHandler) GetAccount(c *gin.Context) {
	accountIdP, err := strconv.Atoi(c.Param("id"))
	accountId := int64(accountIdP)

	if err != nil {
		c.JSON(400, gin.H{"errors": "error id"})
		return
	}
	res, err := h.AccountUseCase.GetById(c.Request.Context(), accountId)
	if err != nil {
		c.JSON(400, gin.H{"errors": err.Error()})
		return
	}
	c.JSON(200, res)
}

func (h *AccountHttpHandler) Login(c *gin.Context) {
}