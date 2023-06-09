package structs

type University struct {
	University_id string `json:"university_id"`
	City          string `json:"city"`
	Country       string `json:"country"`
	Name          string `json:"name"`
}

type Majors struct {
	Major_id      string `json:"major_id"`
	Name          string `json:"name"`
	Category      string `json:"type"`
	University_id string `json:"university_id"`
}

type MajorsWithUniversity struct {
	Major_id        string `json:"major_id"`
	Name            string `json:"name"`
	Category        string `json:"type"`
	University_id   string `json:"university_id"`
	University_name string `json:"university_name"`
}

type MajorName struct {
	Name string `json:"name"`
}

type MajorType struct {
	Type string `json:"type"`
}

type DepartmentWithUniversity struct {
	Department_id   string `json:"department_id"`
	Name            string `json:"name"`
	University_id   string `json:"university_id"`
	University_name string `json:"university_name"`
}

type Department struct {
	Name string `json:"name"`
}

type AddUniversity struct {
	City    string `json:"city"`
	Country string `json:"country"`
	Name    string `json:"name"`
}

type Vote struct {
	Login           string `json:"login"`
	University_name string `json:"university_name"`
}

type MyVotes struct {
	Voted_id        string `json:"voted_id"`
	Login           string `json:"login"`
	University_name string `json:"university_name"`
	Voted_on        string `json:"voted_on"`
}
