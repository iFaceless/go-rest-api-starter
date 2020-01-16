package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ifaceless/go-starter/pkg/web"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Warn(err)
	}
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := web.NewRouter()

	srvPort := os.Getenv("SERVER_PORT")
	if srvPort == "" {
		srvPort = "8001"
	}

	fmt.Printf("Server is listening at: ':%s'\n", srvPort)
	err = r.Run(fmt.Sprintf(":%s", srvPort))
	if err != nil {
		panic(err)
	}
}
