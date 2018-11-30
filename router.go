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
	var user user.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("json: wrong params: %s", err),
		})
		return
	}

	err = r.userService.New(user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("db: query error: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successful",
	})

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

	h := map[string]string{}
	if err := c.ShouldBindJSON(&h); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	u, err := r.userService.Get(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("db: query error: %s", err),
		})
		return
	}

	firstName, foundFirstName := h["first_name"]
	lastName, foundLastName := h["last_name"]
	if !foundFirstName && !foundLastName {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"object":  "error",
			"message": "name not found",
		})
		return
	}

	if foundFirstName && firstName != "" {
		u.FirstName = firstName
	}

	if foundLastName && lastName != "" {
		u.LastName = lastName
	}

	user, err := r.userService.Update(*u)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("db: update error: %s", err),
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
			"message": fmt.Sprintf("db: delete error: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
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

	var account account.Account

	err = c.ShouldBindJSON(&account)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("json: wrong params: %s", err),
		})
		return
	}

	err = r.accountService.New(userID, account)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("db: insert error: %s", err),
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
			"message": fmt.Sprintf("db: delete error: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
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

	h := map[string]string{}
	if err := c.ShouldBindJSON(&h); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	_, ok := h["amount"]
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"object":  "error",
			"message": "amount not found",
		})
		return
	}

	amount, err := strconv.Atoi(h["amount"])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("error: %s", err),
		})
		return
	}

	account, err := r.accountService.Deposit(id, amount)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("db: update error: %s", err),
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

	h := map[string]string{}
	if err := c.ShouldBindJSON(&h); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	_, ok := h["amount"]
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"object":  "error",
			"message": "amount not found",
		})
		return
	}

	amount, err := strconv.Atoi(h["amount"])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("error: %s", err),
		})
		return
	}

	account, err := r.accountService.Withdraw(id, amount)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("db: update error: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, account)
}

func (r *Router) transfer(c *gin.Context) {
	h := map[string]string{}
	if err := c.ShouldBindJSON(&h); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	if _, okAmount := h["amount"]; !okAmount {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"object":  "error",
			"message": "amount not found",
		})
		return
	}

	amount, err := strconv.Atoi(h["amount"])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("error: %s", err),
		})
		return
	}

	if _, okFrom := h["from"]; !okFrom {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"object":  "error",
			"message": "account not found",
		})
		return
	}

	from, err := strconv.Atoi(h["from"])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("error: %s", err),
		})
		return
	}

	if _, okTo := h["to"]; !okTo {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"object":  "error",
			"message": "account not found",
		})
		return
	}

	to, err := strconv.Atoi(h["to"])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("error: %s", err),
		})
		return
	}

	account, err := r.accountService.Transfer(from, to, amount)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("db: update error: %s", err),
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
