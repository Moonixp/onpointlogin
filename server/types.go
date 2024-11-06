package main

import (
	"database/sql"
	"errors"
)

/*
 ---> TABLES
 userids 	-> holds all users and their last loginin time & date
 logintimes -> holds all login times for every user

 linked by id
---> STRUCTS

 -- SqlUser -- is a simple struct that holds data about a user
 that we get from the userids table in the db
 why SqlUser? because of nullvalues in sql queries
 name - the users name
 id - the users id
 Time - the users last login time
 Date - the users last login date

 -- LogintimeUser -- is a simple struct that holds data about a user
 that we get from the logintimes table in the db
 Fullname - the users name
 Time - the users last login time
 Date - the users last login date
 Id - the users id

 -- Userid -- is a simple struct that holds data about a user
 that we get from the userids table in the db
 Fullname - the users name
 LastLoginTime - the users last login time
 LastLoginDate - the users last login date
 Id - the users id
*/

type SqlUser struct {
	name sql.NullString
	id   sql.NullInt64
	Time sql.NullString
	Date sql.NullString
}
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

func (sqlUser *SqlUser) convertoUserid(userid *Userid) error {
	if sqlUser.id.Valid {
		userid.Id = int(sqlUser.id.Int64)
	} else {
		userid.Id = 0
	}
	if sqlUser.name.Valid {
		userid.Fullname = sqlUser.name.String
	} else {
		userid.Fullname = ""
	}
	if sqlUser.Time.Valid {
		userid.LastLoginTime = sqlUser.Time.String
	} else {
		userid.LastLoginTime = ""
	}
	if sqlUser.Date.Valid {
		userid.LastLoginDate = sqlUser.Date.String
	} else {
		userid.LastLoginDate = ""
	}
	return nil
}

func (sqlUser *SqlUser) converttoLoginUser(loginTimeUser *LogintimeUser) error {
	if sqlUser.id.Valid {
		loginTimeUser.Id = int(sqlUser.id.Int64)
	} else {
		loginTimeUser.Id = 0
	}
	if sqlUser.name.Valid {
		loginTimeUser.Fullname = sqlUser.name.String
	} else {
		loginTimeUser.Fullname = ""
	}
	if sqlUser.Time.Valid {
		loginTimeUser.Time = sqlUser.Time.String
	} else {
		loginTimeUser.Time = ""
	}
	if sqlUser.Date.Valid {
		loginTimeUser.Date = sqlUser.Date.String
	} else {
		loginTimeUser.Date = ""
	}
	return nil
}

func (sqlUser *SqlUser) validate() error {
	if sqlUser.name.Valid && sqlUser.id.Valid && sqlUser.Date.Valid && sqlUser.Time.Valid {
		return nil
	} else if !sqlUser.name.Valid {
		return errors.New("name is empty")
	} else if !sqlUser.id.Valid {
		return errors.New("id is empty")
	} else if !sqlUser.Date.Valid {
		return errors.New("date is empty")
	} else if !sqlUser.Time.Valid {
		return errors.New("time is empty")
	}
	return errors.New("unknown error")
}
