package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	T "server/types"

	E "server/error_handler"

	_ "github.com/go-sql-driver/mysql"
)

var DATABASE *sql.DB
var db *sql.DB

func DbInit() {

	db_ip := os.Getenv("DB_IP")
	db_pass := os.Getenv("DB_PASS")

	if db_ip == "" {
		panic("env variable DB_IP is empty")
	}
	if db_pass == "" {
		panic("env variable DB_PASS is empty")
	}

	log.Printf("db_ip : %s", db_ip)
	log.Printf("db_pass : %s", db_pass)

	dsn := fmt.Sprintf("root:%s@tcp(%s:3306)/onpointusers", db_pass, db_ip)

	var err error
	db, err = sql.Open("mysql", dsn)
	E.ErrorHandler(err)

	err = db.Ping()
	E.ErrorHandler(err)

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	log.Println("db connected!")
	DATABASE = db
}

func DbGetAllUsers() ([]T.Userid, error) {
	rows, err := db.Query("SELECT id, fullname, last_login_date, last_login_time FROM users")

	if err != nil {
		return []T.Userid{}, err
	}

	defer rows.Close()

	var UserIds []T.Userid

	for rows.Next() {
		tempuser := T.SqlUser{}
		userid := T.Userid{}

		err := rows.Scan(&tempuser.Id, &tempuser.Name, &tempuser.Date, &tempuser.Time)
		if err != nil {
			return []T.Userid{}, err
		}
		err = tempuser.ConvertoUserid(&userid)
		if err != nil {
			return []T.Userid{}, err
		}
		UserIds = append(UserIds, userid)
	}
	return UserIds, nil
}

func DbGetId(name string) (int, error) {
	if len(name) < 1 {
		return -1, errors.New("name is empty")
	}

	fmt.Println("Name gotten:", name)

	find, err := db.Query("select id, fullname, last_login_date, last_login_time from users where fullname = ?", name)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return -1, err
	}

	tempuser := T.SqlUser{}

	for find.Next() {
		err := find.Scan(&tempuser.Id, &tempuser.Name, &tempuser.Date, &tempuser.Time)

		if err == sql.ErrNoRows {
			return -1, err
		}

		if err != nil {
			return -1, err
		}
	}

	/*
		err = tempuser.Validate()
		if err != nil {
			return -1, err
		}
	*/

	fmt.Println("DbGetID::user:", tempuser)

	if tempuser.Id.Int64 == 0 {
		return -1, errors.New("user not found")
	}

	return int(tempuser.Id.Int64), nil
}

func DbInsertLogin(findUser T.LogintimeUser) error {
	if findUser == (T.LogintimeUser{}) {
		return errors.New("user is empty")
	}

	_, insert_err := db.Exec("INSERT INTO logintimes (id, time, date) VALUES (?, ?, ?)", findUser.Id, findUser.Date, findUser.Time)
	_, insert_err2 := db.Exec("UPDATE users SET last_login_time = ?, last_login_date = ? WHERE fullname = ? AND ID = ?", findUser.Time, findUser.Date, findUser.Fullname, findUser.Id)

	if insert_err != nil {
		return insert_err
	}

	if insert_err2 != nil {
		return insert_err
	}

	return nil
}

func DbisUserLoggedIn(id int) (bool, error) {
	if id < 1 {
		return false, errors.New("id < 1")
	}

	tempuser := T.SqlUser{}

	rows, err := db.Query("SELECT id, fullname, last_login_date, last_login_time FROM users WHERE id = ?", id)

	if err != nil {
		return false, err
	}

	for rows.Next() {
		err := rows.Scan(&tempuser.Id, &tempuser.Name, &tempuser.Date, &tempuser.Time)
		if err != nil {
			return false, err
		}
	}

	fmt.Println("User: ", tempuser)

	if tempuser.Id.Valid && tempuser.Id.Int64 > 0 {
		if tempuser.Date.Valid && tempuser.Date.String == time.Now().Format("2006-01-02") {
			return true, nil
		}
	}

	return false, nil
}
