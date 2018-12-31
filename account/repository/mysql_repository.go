package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/wormcc/marmalade/account"
	"github.com/wormcc/marmalade/models"
)

type mysqlAccountRepository struct {
	Conn *sql.DB
}

func NewMysqlAccountRepository(Conn *sql.DB) account.Repository {

	return &mysqlAccountRepository{Conn}
}
func (r *mysqlAccountRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Account, error) {
	rows, err := r.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]*models.Account, 0)
	for rows.Next() {
		t := new(models.Account)
		err = rows.Scan(
			&t.ID,
			&t.Name,
			&t.Email,
			&t.Password,
			&t.CreateAt,
			&t.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, t)
	}
	return result, nil
}

func (r *mysqlAccountRepository) GetById(ctx context.Context, id int64) (*models.Account, error) {
	query := `SELECT id, name, email, update_at, create_at FROM account WHERE id=?`
	list, err := r.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}
	accountFound := &models.Account{}
	if len(list) > 0 {
		accountFound = list[0]
	} else {
		return nil, models.ErrNotFound
	}
	return accountFound, nil
}

func (r *mysqlAccountRepository) Update(ctx context.Context, account *models.Account) error {
	query := `UPDATE account SET name=?, email=?, password=?, create_at=?, update_at=? WHERE id=?`
	stmt, err := r.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	res, err := stmt.ExecContext(ctx, account.Name, account.Email, account.Password, account.CreateAt, account.UpdateAt, account.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		err = fmt.Errorf("sql-error:Weird  Behaviour. Total Affected: %d", rowsAffected)
		return err
	}
	return nil
}

func (r *mysqlAccountRepository) Store(ctx context.Context, account *models.Account) error {
	return nil
}
