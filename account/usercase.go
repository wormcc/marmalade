package account

import (
	"context"
	"github.com/wormcc/marmalade/models"
)

type UseCase interface {
	GetById(ctx context.Context, id int64) (*models.Account, error)
	Store(ctx context.Context, account *models.Account) error
	UpdatePassword(ctx context.Context, password string, accountId int64) error
}
