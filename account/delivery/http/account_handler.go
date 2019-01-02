package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/wormcc/marmalade/account"
	"github.com/wormcc/marmalade/common"
	"github.com/wormcc/marmalade/middleware"
	"net/http"
)

type AccountHttpHandler struct {
	AccountUseCase account.UseCase
}

func NewAccountHttpHandler(r *gin.Engine, us account.UseCase) {
	handler := &AccountHttpHandler{
		AccountUseCase: us,
	}
	accountGroup := r.Group("/users")
	accountGroup.Use(middleware.AuthMiddleware(true))
	{
		accountGroup.GET("/whoami", handler.GetAccount)
		accountGroup.PUT("/password", handler.UpdatePassword)
	}
	authGroup := r.Group("/auth")
	authGroup.Use(middleware.AuthMiddleware(false))
	{
		authGroup.POST("/register", handler.CreateAccount)
		authGroup.POST("/login", handler.Login)
	}
}

func (h *AccountHttpHandler) CreateAccount(c *gin.Context) {
	accountModelValidator := NewAccountModelValidator()
	if err := accountModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	accountModel := accountModelValidator.accountModel
	if err := h.AccountUseCase.Store(c, &accountModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Created!"})
}

func (h *AccountHttpHandler) GetAccount(c *gin.Context) {
	currentAccountId := c.MustGet("current_account_id").(int64)

	res, err := h.AccountUseCase.GetById(c.Request.Context(), currentAccountId)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"errors": err.Error()})
		return
	}
	accountSerializer := AccountSerializer{c, res}
	c.JSON(200, accountSerializer.Response())
}

func (h *AccountHttpHandler) UpdatePassword(c *gin.Context) {
	accountId := c.MustGet("current_account_id").(int64)
	accountModel, err := h.AccountUseCase.GetById(c.Request.Context(), accountId)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	accountUpdatePasswordValidator := NewAccountUpdatePasswordValidator(accountModel)
	if err := accountUpdatePasswordValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	accountUpdatePasswordModel := accountUpdatePasswordValidator.accountModel
	if err := h.AccountUseCase.Update(c.Request.Context(), &accountUpdatePasswordModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (h *AccountHttpHandler) Login(c *gin.Context) {
	loginValidator := NewLoginValidator()
	if err := loginValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	accountModel, err := h.AccountUseCase.GetByEmail(c.Request.Context(), loginValidator.Email)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("login", err))
		return
	}
	if !common.CheckPassword(loginValidator.Password, accountModel.Password) {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("login", errors.New("invalid email or password")))
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": common.GenToken(accountModel.ID)})
}
