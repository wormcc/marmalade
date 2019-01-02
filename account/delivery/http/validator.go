package http

import (
	"github.com/gin-gonic/gin"
	"github.com/wormcc/marmalade/common"
	"github.com/wormcc/marmalade/models"
	"time"
)

type AccountModelValidator struct {
	Account struct {
		Name            string `form:"name" json:"name" binding:"exists,alphanum,min=4,max=100"`
		Email           string `form:"email" json:"email" binding:"exists,email"`
		Password        string `form:"password" json:"password" binding:"exists,min=8,max=255"`
		PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"exists,eqfield=Password"`
	} `json:"account"`
	accountModel models.Account
}

func (validator *AccountModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, validator)
	if err != nil {
		return err
	}
	validator.accountModel.Email = validator.Account.Email
	validator.accountModel.Name = validator.Account.Name
	validator.accountModel.Password = common.SetPassword(validator.Account.Password)
	validator.accountModel.CreateAt = time.Now()
	validator.accountModel.UpdateAt = time.Now()
	return nil
}

func NewAccountModelValidator() AccountModelValidator {
	accountModelValidator := AccountModelValidator{}
	return accountModelValidator
}

type AccountUpdatePasswordValidator struct {
	Password        string `form:"password" json:"password" binding:"exists,min=8,max=255"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"exists,eqfield=Password"`

	accountModel models.Account
}

func (validator *AccountUpdatePasswordValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, validator)
	if err != nil {
		return err
	}
	validator.accountModel.Password = common.SetPassword(validator.Password)
	validator.accountModel.UpdateAt = time.Now()
	return nil
}

func NewAccountUpdatePasswordValidator(accountModel *models.Account) AccountUpdatePasswordValidator {
	accountUpdatePasswordValidator := AccountUpdatePasswordValidator{}
	accountUpdatePasswordValidator.accountModel = *accountModel
	return accountUpdatePasswordValidator
}

type LoginValidator struct {
	Password string `form:"password" json:"password" binding:"exists,min=8,max=255"`
	Email    string `form:"email" json:"email" binding:"exists,email"`

	accountModel models.Account
}

func (validator *LoginValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, validator)
	if err != nil {
		return err
	}
	return nil
}

func NewLoginValidator() LoginValidator {
	loginValidator := LoginValidator{}
	return loginValidator
}