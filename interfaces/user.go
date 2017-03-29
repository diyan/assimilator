package interfaces

// User is an interface which describes the authenticated User for a request.
//
// You should provide **at least** either an `id` (a unique identifier for
// an authenticated user) or `ip_address` (their IP address).
//
// All other attributes are optional.
//
// {
//     "id": "unique_id",
//     "username": "my_user",
//     "email": "foo@example.com"
//     "ip_address": "127.0.0.1",
//     "optional": "value"
// }
type User struct {
	ID        string `in:"id"         json:"id"`
	Username  string `in:"username"   json:"username"`
	Email     string `in:"email"      json:"email"`
	IPAddress string `in:"ip_address" json:"ip_address"`
	// TODO Does Sentry allows arbitrary key/value pairs for User interface?
	Extra map[string]string `in:"-"   json:"-"`
}

func init() {
	Register(&User{})
}

func (*User) KeyAlias() string {
	return "user"
}

func (*User) KeyCanonical() string {
	return "sentry.interfaces.User"
}

func (user *User) DecodeRequest(request map[string]interface{}) error {
	// TODO process extra fields
	return nil
}
