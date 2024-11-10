package main

import (
	"database/sql"
	"server/database"
	"server/login_handler"
	"server/user_handler"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB = database.DATABASE

func main() {

	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Use(CORSMiddleware())
	database.DbInit()

	router.GET("/users", user_handler.GetAllUsers)
	router.POST("/loginuser", login_handler.Loginbybody)
	router.GET("/login", login_handler.Loginbyparam)
	router.GET("/names", user_handler.GetUsersName)
	router.GET("/loggedin", login_handler.IsUserLoggedIn)
	router.Run(":8000")
	db.Close()

}
