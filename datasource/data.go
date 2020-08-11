package datasource

var AllUsers = map[int]User{}

type User struct {
	ID          int    `json:"id"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
}
