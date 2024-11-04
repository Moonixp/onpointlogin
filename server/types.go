package main

/*
2 TABLES
userids => holds all users and their last loginin time & date
logintimes => holds all login times for everyuser

linked by  id
hence 2 user types describing them below
*/
type LogintimeUser struct {
	Fullname string `json:"fullname"`
	Time     string `json:"time"`
	Date     string `json:"date"`
	Id       int    `json:"id"`
}

type Userid struct {
	Fullname      string `json:"fullname"`
	LastLoginTime string `json:"last_login_time"`
	LastLoginDate string `json:"last_login_date"`
	Id            int    `json:"id"`
}
