package structs

type User struct {
	User_id  string `json:"user_id"`
	Email    string `json:"email"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
