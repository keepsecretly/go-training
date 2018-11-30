package main

import (
	"go-training/user"
	"os"
)

func main() {
	r := &Router{
		userService: &user.UserServiceImp{},
	}

	ge := setupRouter(r)
	ge.Run(":" + os.Getenv("PORT"))

}
