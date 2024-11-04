package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func SetDateAndTime(user *LogintimeUser) error {
	if user == nil {
		return errors.New("user is nil")
	}
	now := time.Now()
	user.Time = now.Format("15:04")
	user.Date = now.Format("2006-01-02")
	return nil
}

func GetAllUsers(ctx *gin.Context) {
	users, err := DbGetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func GetUsersName(ctx *gin.Context) {
	names, err := os.ReadFile("names.txt")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	lines := strings.Split(string(names), "\n")

	ctx.JSON(http.StatusOK, lines)
}

func Loginbyparam(ctx *gin.Context) {

	name := ctx.Query("name")
	fmt.Println("name: ", name)

	alreadyLoggedIn, id, err := DbGetId(name)
	println("id: ", id, " err: ", err)

	if (err != nil) || (id == -1) || (id == 0) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	var findUser LogintimeUser = LogintimeUser{
		Id:       id,
		Fullname: name,
	}

	if err := SetDateAndTime(&findUser); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("loginByParam User:  ", findUser)

	if alreadyLoggedIn {
		ctx.JSON(200, gin.H{"id": -1, "data": "user already logged in"})
		return
	}

	err = DbInsertLogin(findUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, findUser)
}

func Loginbybody(ctx *gin.Context) {
	type objdata struct {
		Name string `json:"name"`
	}
	obj := objdata{}

	if err := ctx.ShouldBind(&obj); err != nil {
		fmt.Println("bad data : ", obj, " err: ", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	alreadyLoggedIn, id, err := DbGetId(obj.Name)
	println("id: ", id, " err: ", err)

	if err != nil || id == -1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	user := LogintimeUser{
		Id:       id,
		Fullname: obj.Name,
	}

	if err := SetDateAndTime(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("LoginbyBody User", user)
	fmt.Println("is already logged in: ", alreadyLoggedIn)

	if alreadyLoggedIn {
		ctx.JSON(200, gin.H{"id": -1, "data": "user already logged in"})
		return
	}

	err = DbInsertLogin(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, user)
}
