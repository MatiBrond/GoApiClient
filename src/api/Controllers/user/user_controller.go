package user

import (
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/taller-go/src/api/Services/user"
	"github.com/mercadolibre/taller-go/src/api/Utils"
	"net/http"
	"strconv"
)

const(

	paramUserID = "id"
)

func GetUserByIdS(context *gin.Context){

	id := context.Param(paramUserID)
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		apiError := &Utils.ApiError{
			Message: "Fatal URL",
			Status: http.StatusBadRequest}
		context.JSON(apiError.Status, apiError)
		return
	}
		user, apiError := user.GetUserByIdSer(userID);
		if apiError != nil{
			context.JSON(apiError.Status, apiError)
			return
	}
	context.JSON(200, user)
}

func Mostrar(context *gin.Context){

	id := context.Param(paramUserID)
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		apiError := &Utils.ApiError{
			Message: "Fatal URL",
			Status: http.StatusBadRequest}
		context.JSON(apiError.Status, apiError)
		return
	}
	user, apiError := user.GetUserByIdSe(userID);
	if apiError != nil{
		context.JSON(apiError.Status, apiError)
		return
	}
	context.JSON(200, user)
}

