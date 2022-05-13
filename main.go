package main

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"jwt-practice/database"
	"jwt-practice/router"
)

func main()  {
	flag.Parsed()

	err := database.Init()
	if err != nil {
		fmt.Println("db conn err")
		return
	}
	logrus.Info("website run")
	router.InitRouter()

}
