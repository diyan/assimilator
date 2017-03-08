package interfaces

import "fmt"

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
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	IPAddress string `json:"ip_address"`
	// Is Sentry allows arbitrary key/value pairs for User interface?
	Extra map[string]string `json:"-"`
}

func (user *User) UnmarshalRecord(nodeBlob interface{}) error {
	return nil
}

func (user *User) UnmarshalAPI(rawEvent map[string]interface{}) error {
	if rawUser, ok := rawEvent["user"].(map[string]interface{}); ok {
		// TODO validate input
		user.ID = anyTypeToString(rawUser["id"])
		user.Username = anyTypeToString(rawUser["username"])
		user.Email = anyTypeToString(rawUser["email"])
		user.IPAddress = anyTypeToString(rawUser["ip_address"])
		// TODO process extra fields
	}
	return nil
}

// TODO duplicated code in interfaces and web/api/store_event.go
func anyTypeToString(v interface{}) string {
	if v != nil {
		return fmt.Sprint(v)
	}
	return ""
}
