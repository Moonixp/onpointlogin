package user_handler

import (
	"net/http"
	"os"
	"server/database"
	"strings"

	"errors"
	T "server/types"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func SetDateAndTime(user *T.LogintimeUser) error {
	if user == nil {
		return errors.New("user is nil")
	}
	now := time.Now()
	user.Time = now.Format("15:04")
	user.Date = now.Format("2006-01-02")
	return nil
}

func GetAllUsers(ctx *gin.Context) {
	users, err := database.DbGetAllUsers()
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
