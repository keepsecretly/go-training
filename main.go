package main

import (
	"go-training/account"
	"go-training/user"
	"log"
	"os"

	mgo "gopkg.in/mgo.v2"
)

func main() {
	sess, err := mgo.Dial(os.Getenv("MLAB_URI"))
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	r := &Router{
		userService: &user.UserServiceImp{
			Session: sess,
		},
		accountService: &account.AccountServiceImp{
			Session: sess,
		},
		adminService: &AdminServiceImp{
			Session: sess,
		},
	}

	ge := setupRouter(r)
	ge.Run(":" + os.Getenv("PORT"))

}
