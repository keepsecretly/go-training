package main

import (
	"fmt"
	"go-training/account"
	"go-training/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Router struct {
	userService    user.UserService
	accountService account.AccountService
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

func (r *Router) addAccount(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("error: %s", err),
		})
		return
	}

	account, err := r.accountService.New(userID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("db: query error: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, account)
}

func (r *Router) allAccounts(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("error: %s", err),
		})
		return
	}

	accounts, err := r.accountService.All(userID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("db: query error: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, accounts)
}

func (r *Router) deleteAccount(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("error: %s", err),
		})
		return
	}

	err = r.accountService.Delete(id)

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

func (r *Router) deposit(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("error: %s", err),
		})
		return
	}

	account, err := r.accountService.Deposit(id, 0)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("db: query error: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, account)
}

func (r *Router) withdraw(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("error: %s", err),
		})
		return
	}

	account, err := r.accountService.Withdraw(id, 0)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("db: query error: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, account)
}

func setupRouter(r *Router) *gin.Engine {
	g := gin.Default()

	u := g.Group("/users")
	b := g.Group("/bankAccounts")

	u.GET("/", r.allUser)
	u.POST("/", r.addUser)

	u.GET("/:id", r.getUserByID)
	u.PUT("/:id", r.updateUser)
	u.POST("/:id", r.updateUser)
	u.DELETE("/:id", r.deleteUser)

	u.GET("/:id/bankAccounts/", r.allAccounts)
	u.POST("/:id/bankAccounts/", r.addAccount)

	b.DELETE("/:id", r.deleteAccount)
	b.PUT("/:id/deposit", r.deposit)
	b.POST("/:id/deposit", r.deposit)
	b.PUT("/:id/withdraw", r.withdraw)
	b.POST("/:id/withdraw", r.withdraw)

	return g
}
