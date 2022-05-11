package main

import "jwt-practice/util"

func main() {
	id, _ := util.GenerateUUid("user")
	print(id)
}
