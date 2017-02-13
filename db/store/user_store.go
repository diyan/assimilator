package store

import (
	"github.com/diyan/assimilator/db"
	"github.com/diyan/assimilator/models"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

type UserStore struct {
	c echo.Context
}

func NewUserStore(c echo.Context) UserStore {
	return UserStore{c: c}
}

func (s UserStore) SaveUser(user models.User) error {
	db, err := db.FromE(s.c)
	if err != nil {
		return errors.Wrap(err, "failed to save user")
	}
	_, err = db.InsertInto("auth_user").
		Columns("id", "password", "last_login", "username", "first_name", "email", "is_staff", "is_active", "is_superuser", "date_joined", "is_managed", "is_password_expired", "last_password_change", "session_nonce").
		Record(user).
		Exec()
	return errors.Wrap(err, "failed to save user")
}
