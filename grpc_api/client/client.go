package main

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc/proto"
	"log"
	"net/http"
)

var serverAddress = "localhost:50051"

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewLoginClient(conn)

	r := gin.Default()
	r.GET("/User/:username/:password", func(c *gin.Context) {
		username := c.Param("username")
		password := c.Param("password")
		log.Print(username, password)
		res, err := client.Login(c, &pb.LoginRequest{Username: username, Password: password})
		log.Println(res)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": res.Result,
		})
	})
	r.Run()
}
