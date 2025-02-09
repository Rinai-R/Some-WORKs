package main

import (
	user "Golang/2025/02February/20250209/kitex/kitex_gen/user/user"
	"log"
)

func main() {
	svr := user.NewServer(new(UserImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
