package models

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	CareerId int    `json:"career_id"`
	TypeId   int    `json:"type_id"`
	Status   bool   `json:"-"`
}
