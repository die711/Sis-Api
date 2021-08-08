package models

type Course struct {
	Id       int    `json:"id"`
	CareerId int    `json:"career_id"`
	Name     string `json:"name"`
	Credits  int    `json:"credits"`
	Status   bool   `json:"-"`
}
