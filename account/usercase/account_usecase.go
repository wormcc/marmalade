package usercase

import (
	"context"
	"github.com/wormcc/marmalade/account"
	"github.com/wormcc/marmalade/models"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type accountUseCase struct {
	accountRepo    account.Repository
	contextTimeout time.Duration
}

func NewAccountUseCase(accountRepo account.Repository, timeout time.Duration) account.UseCase {
	return &accountUseCase{accountRepo: accountRepo, contextTimeout: timeout}
}

func (au *accountUseCase) GetById(c context.Context, accountId int64) (*models.Account, error) {
	ctx, cancel := context.WithTimeout(c, au.contextTimeout)
	defer cancel()
	res, err := au.accountRepo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (au *accountUseCase) Store(c context.Context, account *models.Account) error {
	return nil
}

func (au *accountUseCase) UpdatePassword(c context.Context, password string, accountId int64) error {
	ctx, cancel := context.WithTimeout(c, au.contextTimeout)
	defer cancel()
	res, err := au.accountRepo.GetById(ctx, accountId)
	if err != nil {
		return err
	}
	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	res.Password = string(passwordHash)
	err = au.accountRepo.Update(ctx, res)
	return err
}