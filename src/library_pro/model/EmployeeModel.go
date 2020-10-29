package model

type Employee struct {
	Id int `form:"id" json:"id" xml:"id"`
	Chinese_name string `form:"chinese_name" json:"chinese_name" xml:"chinese_name"`
	English_name      string  `json:"english_name" form:"english_name"`
	Position_name  string `json:"position_name" form:"position_name"`
	Birthday  string `json:"birthday" form:"birthday"`
	Status  int `json:"status" form:"status"`
	Create_time string `json:"create_time" form:"create_time"`
}
