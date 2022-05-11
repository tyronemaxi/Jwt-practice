package main

import (
	"flag"
	"fmt"
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
	router.InitRouter()
}
