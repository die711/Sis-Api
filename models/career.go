package models

type Career struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Status bool   `json:"-"`
}
