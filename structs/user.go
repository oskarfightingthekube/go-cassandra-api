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
	Login    string `json:"login"`
	Password string `json:"password"`
}

type University struct {
	University_id string `json:"university_id"`
	City          string `json:"city"`
	Country       string `json:"country"`
	Name          string `json:"name"`
}
