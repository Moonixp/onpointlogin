package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {

	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Use(CORSMiddleware())
	DbInit()

	router.GET("/users", GetAllUsers)
	router.POST("/loginuser", Loginbybody)
	router.GET("/login", Loginbyparam)
	router.GET("/names", GetUsersName)
	router.GET("/loggedin", isUserLoggedIn)
	router.Run(":8000")
	db.Close()

}
