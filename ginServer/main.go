package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Return a default engine with attached recovery and logger
	r := gin.Default()

	// Running server on 8080 by default
	//r.Run()
	r.GET("/testServer", TestServer)
	r.GET("/geUserData", GetUserData)
	// Given hardcoded port - server will run on 3030
	r.Run(":3030")
}

// TestServer ... Check if server is running
func TestServer(c *gin.Context) {
	// Simple JSON request
	c.JSON(200, gin.H{"message": "Gin server running on 3030"})
}

//GetUserData ... Get basic user data using struct
func GetUserData(c *gin.Context) {
	var UserData struct {
		Name  string `json:"user"`
		Age   int
		Place string
	}
	UserData.Name = "ABC"
	UserData.Age = 25
	UserData.Place = "XYZ"
	c.JSON(200, UserData)
}
