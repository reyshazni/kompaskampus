package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

var users = []user{
	{Id: 13, Username: "Yer"},
	{Id: 15, Username: "Andicka"},
	{Id: 58, Username: "Yandick"},
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func main() {
	router := gin.Default()

	// new `GET /users` route associated with our `getUsers` function
	router.GET("/users", getUsers)

	router.Run("localhost:8080")
}
