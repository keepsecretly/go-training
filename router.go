package main

import (
	"fmt"
	"go-training/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Router struct {
	userService user.UserService
}

func (r *Router) allUser(c *gin.Context) {
	users, err := r.userService.All()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("db: query error: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (r *Router) getUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("error: %s", err),
		})
		return
	}

	user, err := r.userService.Get(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("db: query error: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (r *Router) addUser(c *gin.Context) {
	user, err := r.userService.New()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("db: query error: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (r *Router) updateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("error: %s", err),
		})
		return
	}

	user, err := r.userService.Update(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("db: query error: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (r *Router) deleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("error: %s", err),
		})
		return
	}

	err = r.userService.Delete(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("db: query error: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"object":  "error",
		"message": "Successful",
	})
}

func setupRouter(r *Router) *gin.Engine {
	g := gin.Default()

	u := g.Group("/users")

	u.GET("/", r.allUser)
	u.POST("/", r.addUser)

	u.GET("/:id", r.getUserByID)
	u.PUT("/:id", r.updateUser)
	u.POST("/:id", r.updateUser)
	u.DELETE("/:id", r.deleteUser)

	return g
}
