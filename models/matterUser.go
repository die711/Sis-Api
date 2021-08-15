package models

type MatterUser struct {
	UserId   int `json:"user_id"`
	MatterId int `json:"matter_id"`
	Cal1     int `json:"cal1"`
	Cal2     int `json:"cal2"`
	Cal3     int `json:"cal3"`
	Cal4     int `json:"cal4"`
}
