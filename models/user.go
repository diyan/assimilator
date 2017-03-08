package models

import "time"

type User struct {
	ID       int    `db:"id" json:"id,string"`
	Username string `db:"username" json:"username"`
	// This column is called first_name for legacy reasons, but it is the entire
	// display name
	Name               string    `db:"first_name" json:"name"`
	Email              string    `db:"email" json:"email"`
	IsStaff            bool      `db:"is_staff" json:"-"`
	IsActive           bool      `db:"is_active" json:"-"`
	IsSuperuser        bool      `db:"is_superuser" json:"-"`
	IsManaged          bool      `db:"is_managed" json:"-"`
	IsPasswordExpired  bool      `db:"is_password_expired" json:"-"`
	LastPasswordChange time.Time `db:"last_password_change" json:"-"`
	Password           string    `db:"password" json:"-"`
	LastLogin          time.Time `db:"last_login" json:"-"`
	DateCreated        time.Time `db:"date_joined" json:"-"`
	// TODO add session_nonce *string
}

func (user *User) PostGet() {
	if user.Name != "" {
		return
	}
	if user.Email != "" {
		user.Name = user.Email
	} else {
		user.Name = user.Username
	}
}
