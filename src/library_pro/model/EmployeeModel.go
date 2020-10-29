package model

type Employee struct {
	Id       uint16 `json:"id" form:"id"`
	Chinese_name string `json:"chinese_name" form:"chinese_name"`
	English_name      string  `json:"english_name" form:"english_name"`
	Position_name  string `json:"position_name" form:"position_name"`
	Birthday  string `json:"birthday" form:"birthday"`
	Status  int `json:"status" form:"status"`
	Create_time string `json:"create_time" form:"create_time"`
}
