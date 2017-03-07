package interfaces

type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	IPAddress string `json:"ip_address"`
}
