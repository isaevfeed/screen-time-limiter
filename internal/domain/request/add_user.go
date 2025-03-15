package request

type AddUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
