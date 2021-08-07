package models

type Student struct {
	NoControl int    `json:"no_control"`
	Name      string `json:"name"`
	LastName  string `json:"last_name"`
	Career    string `json:"career"`
	Semester  int    `json:"semester"`
}


