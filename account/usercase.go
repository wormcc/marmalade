package account

import (
	"context"
	"github.com/wormcc/marmalade/models"
)

type UseCase interface {
	GetById(ctx context.Context, id int64) (*models.Account, error)
	GetByEmail(ctx context.Context, email string) (*models.Account, error)
	Store(ctx context.Context, account *models.Account) error
	Update(ctx context.Context, accountModel *models.Account) error
}
