package main

import (
	"go-training/user"

	"github.com/gin-gonic/gin"
)

type Router struct {
	userService user.UserService
}

func (r *Router) allUser(c *gin.Context) {
	r.userService.All()
}

func setupRouter(r Router) *gin.Engine {
	g := gin.Default()

	u := g.Group("/users")

	u.GET("/", r.allUser)

	return g
}
