package login_handler

import (
	"fmt"
	"net/http"
	"strconv"

	"server/database"
	T "server/types"
	"server/user_handler"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Login takes a gin context, a name and an id.
//
// It checks if the user is already logged in, if so it returns a 200 with a json object
// containing the id of the user and a "data" field saying "user already logged in".
//
// If the user is not logged in, it will log the user in and then return a 200 with a json
// object containing the user's name, id and the current time and date.
//
// If there is an error when logging the user in, it will return a 500 with a json object
// containing the error message.
func Login(ctx *gin.Context, name string, id int) (T.LogintimeUser, error) {

	alreadyLoggedIn, err := database.DbisUserLoggedIn(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return T.LogintimeUser{}, err
	}

	findUser := T.LogintimeUser{
		Id:       id,
		Fullname: name,
	}

	if err := user_handler.SetDateAndTime(&findUser); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return T.LogintimeUser{}, err
	}

	fmt.Println("loginByParam::user: ", findUser)

	if alreadyLoggedIn {
		ctx.JSON(200, gin.H{"id": -1, "data": "user already logged in"})
		return T.LogintimeUser{}, err
	}

	err = database.DbInsertLogin(findUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return T.LogintimeUser{}, err
	}
	return findUser, nil
}

// Loginbyparam takes a query param of name and logs that user in.
// then implements func Login
func Loginbyparam(ctx *gin.Context) {
	name := ctx.Query("name")
	fmt.Println("LoginbyParam::name: ", name)

	id, err := database.DbGetId(name)
	if (err != nil) || (id == -1) || (id == 0) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	user, err := Login(ctx, name, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, user)
}

// Loginbybody takes a json body of {name: string} and logs that user in.
// then implements func Login
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

	id, err := database.DbGetId(obj.Name)
	println("id: ", id, " err: ", err)

	if err != nil || id == -1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	user, err := Login(ctx, obj.Name, id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, user)
}

func IsUserLoggedIn(ctx *gin.Context) {
	id := ctx.Query("id")
	fmt.Println("id: ", id)

	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id must be an integer"})
		return
	}

	value, err := database.DbisUserLoggedIn(int(idInt))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if value {
		ctx.JSON(http.StatusOK, gin.H{"data": "user is logged in"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "user is not logged in"})
}
