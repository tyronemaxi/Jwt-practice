package main

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"website/database"
	"website/router"
)

func main() {
	flag.Parsed()

	err := database.Init()
	if err != nil {
		fmt.Println("db conn err")
		return
	}
	logrus.Info("website run")
	router.InitRouter()

}
