package models

type Matter struct {
	Id       int    `json:"id"`
	UserId   string `json:"user_id"`
	CourseId string `json:"course_id"`
}
