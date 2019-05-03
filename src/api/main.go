package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/taller-go/src/api/Controllers/ping"

	"github.com/mercadolibre/taller-go/src/api/Controllers/user"
)


const (
	port = ":8080"
)

var(
 router = gin.Default()
 )



func main(){

	router.GET("/ping", ping.Ping)
	router.GET("/users/:id", user.GetUserByIdS)
	router.GET("/mostrar", user.Mostrar)

	router.Run(port)
}