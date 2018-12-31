package account

import (
	"context"
	"github.com/wormcc/marmalade/models"
)

type Repository interface {
	GetById(ctx context.Context, id int64) (*models.Account, error)
	Store(ctx context.Context, account *models.Account) error
	Update(ctx context.Context, account *models.Account) error
}
