package ping

import "github.com/gin-gonic/gin"


func Ping (context *gin.Context){



	context.String(200, "pong")

}


