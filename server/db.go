package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func DbInit() {
	dsn := "root:root@tcp(localhost:3306)/onpointusers"
	var err error

	db, err = sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	fmt.Println("db connected!")
}

func DbGetAllUsers() ([]Userid, error) {
	rows, err := db.Query("SELECT fullname FROM userids")

	if err != nil {
		return []Userid{}, err
	}

	defer rows.Close()

	var UserIds []Userid

	for rows.Next() {
		var userid Userid

		if err := rows.Scan(&userid.Fullname); err != nil {
			return []Userid{}, err
		}

		UserIds = append(UserIds, userid)
	}
	return UserIds, nil
}

func DbGetId(name string) (bool, int, error) {
	if len(name) < 1 {
		return false, -1, errors.New("name is empty")
	}
	fmt.Println("Name gotten:", name)

	find, err := db.Query("select id, fullname, last_login_date, last_login_time from userids where fullname = ?", name)
	if err != nil {
		return false, -1, err
	}
	user := Userid{}
	var lastLoginDate, lastLoginTime sql.NullString

	for find.Next() {
		err := find.Scan(&user.Id, &user.Fullname, &lastLoginDate, &lastLoginTime)

		if err == sql.ErrNoRows {
			return false, -1, err
		}

		if err != nil {
			return false, -1, err
		}

	}

	if lastLoginDate.Valid {
		user.LastLoginDate = lastLoginDate.String
	} else {
		user.LastLoginDate = ""
	}

	if lastLoginTime.Valid {
		user.LastLoginTime = lastLoginTime.String
	} else {
		user.LastLoginTime = ""
	}

	fmt.Println("DbGetID user:", user)

	if user.Id == 0 {
		return false, -1, nil
	}

	if user.LastLoginDate == time.Now().Format("2006-01-02") {
		return true, user.Id, nil
	}

	return false, user.Id, nil
}

func DbInsertLogin(findUser LogintimeUser) error {
	if findUser == (LogintimeUser{}) {
		return errors.New("user is empty")
	}

	_, insert_err := db.Exec("INSERT INTO logintimes (id, time, date) VALUES (?, ?, ?)", findUser.Id, findUser.Date, findUser.Time)
	_, insert_err2 := db.Exec("UPDATE userids SET last_login_time = ?, last_login_date = ? WHERE fullname = ? AND ID = ?", findUser.Time, findUser.Date, findUser.Fullname, findUser.Id)

	if insert_err != nil {
		return insert_err
	}

	if insert_err2 != nil {
		return insert_err
	}

	return nil
}
