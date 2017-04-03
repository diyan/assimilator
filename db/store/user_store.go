package store

import (
	"github.com/diyan/assimilator/models"
	"github.com/gocraft/dbr"
	"github.com/pkg/errors"
)

type UserStore struct {
}

func NewUserStore() UserStore {
	return UserStore{}
}

func (s UserStore) SaveUser(tx *dbr.Tx, user models.User) error {
	_, err := tx.InsertInto("auth_user").
		Columns("id", "password", "last_login", "username", "first_name", "email", "is_staff", "is_active", "is_superuser", "date_joined", "is_managed", "is_password_expired", "last_password_change", "session_nonce").
		Record(user).
		Exec()
	return errors.Wrap(err, "failed to save user")
}
