package repository_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	accountRepo "github.com/wormcc/marmalade/account/repository"
	"github.com/wormcc/marmalade/models"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"time"
)

func TestMysqlAccountRepository_GetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "update_at", "create_at"}).
		AddRow(1, "test", "test@test.com", "test_password", time.Now(), time.Now())
	query := "^SELECT \\* FROM account WHERE id=\\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := accountRepo.NewMysqlAccountRepository(db)

	num := int64(5)
	anAccount, err := a.GetById(context.TODO(), num)
	assert.NoError(t, err)
	assert.NotNil(t, anAccount)
}

func TestMysqlAccountRepository_GetByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "update_at", "create_at"}).
		AddRow(1, "test", "test@test.com", "test_password", time.Now(), time.Now())
	query := "^SELECT \\* FROM account WHERE email=\\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := accountRepo.NewMysqlAccountRepository(db)

	email := "test@test.com"
	anAccount, err := a.GetByEmail(context.TODO(), email)
	assert.NoError(t, err)
	assert.NotNil(t, anAccount)
}

func TestMysqlAccountRepository_Store(t *testing.T) {
	now := time.Now()
	ar := &models.Account{
		Name:     "test",
		Email:    "test@test.com",
		Password: "test_password",
		CreateAt: now,
		UpdateAt: now,
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	query := "INSERT account SET name=\\?, email=\\?, password=\\?, create_at=\\?, update_at=\\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(ar.Name, ar.Email, ar.Password, ar.CreateAt, ar.UpdateAt).WillReturnResult(sqlmock.NewResult(12, 1))

	a := accountRepo.NewMysqlAccountRepository(db)

	err = a.Store(context.TODO(), ar)
	assert.NoError(t, err)
	assert.Equal(t, int64(12), ar.ID)
}

func TestMysqlAccountRepository_Update(t *testing.T) {
	now := time.Now()
	ar := &models.Account{
		ID:       12,
		Name:     "test",
		Email:    "test@test.com",
		Password: "test_password",
		CreateAt: now,
		UpdateAt: now,
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	query := "UPDATE account SET name=\\?, email=\\?, password=\\?, create_at=\\?, update_at=\\? WHERE id=\\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(ar.Name, ar.Email, ar.Password, ar.CreateAt, ar.UpdateAt, ar.ID).WillReturnResult(sqlmock.NewResult(12, 1))

	a := accountRepo.NewMysqlAccountRepository(db)

	err = a.Update(context.TODO(), ar)
	assert.NoError(t, err)
}
