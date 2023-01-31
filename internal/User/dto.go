package User

type CreateUser struct {
	id       string `json:"id"`
	name     string `json:"name"`
	surname  string `json:"surname"`
	password string `json:"password"`
	role     int    `json:"role"`
}
