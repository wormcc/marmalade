package usercase

import (
	"context"
	"fmt"
	"github.com/wormcc/marmalade/account"
	"github.com/wormcc/marmalade/models"
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
	res, err := au.accountRepo.GetById(ctx, accountId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (au *accountUseCase) GetByEmail(c context.Context, email string) (*models.Account, error) {
	ctx, cancel := context.WithTimeout(c, au.contextTimeout)
	defer cancel()
	res, err := au.accountRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (au *accountUseCase) Store(c context.Context, account *models.Account) error {
	ctx, cancel := context.WithTimeout(c, au.contextTimeout)
	defer cancel()
	existedAccount, err := au.accountRepo.GetByEmail(ctx, account.Email)
	fmt.Println(existedAccount, err)
	if existedAccount != nil {
		return models.ErrUnique
	}
	err = au.accountRepo.Store(ctx, account)
	if err != nil {
		return err
	}
	return nil
}

func (au *accountUseCase) Update(c context.Context, accountModel *models.Account) error {
	ctx, cancel := context.WithTimeout(c, au.contextTimeout)
	defer cancel()
	err := au.accountRepo.Update(ctx, accountModel)
	return err
}
