package structs

type User struct {
	User_id  string `json:"user_id"`
	Email    string `json:"email"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type AddUser struct {
	Email    string `json:"email"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginUser struct {
	Login    string `json:"email"`
	Password string `json:"password"`
}
type LoginUser2 struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
