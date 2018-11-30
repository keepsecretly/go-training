package main

import (
	"go-training/account"
	"go-training/user"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testUserServiceImp struct {
}

type testAccountServiceImp struct {
}

func (s *testAccountServiceImp) All(user int) ([]account.Account, error) {
	accs := []account.Account{}
	return accs, nil
}

func (s *testAccountServiceImp) New(user int, account account.Account) error {
	return nil
}

func (s *testAccountServiceImp) Deposit(id int, amount int) (*account.Account, error) {
	return &account.Account{}, nil
}

func (s *testAccountServiceImp) Withdraw(id int, amount int) (*account.Account, error) {
	return &account.Account{}, nil
}

func (s *testAccountServiceImp) Delete(id int) error {
	return nil
}

func (s *testAccountServiceImp) Transfer(fromID int, toID int) (*account.Account, error) {
	return &account.Account{}, nil
}

func (s *testUserServiceImp) All() ([]user.User, error) {
	users := []user.User{
		{
			ID:        1,
			FirstName: "NiM",
			LastName:  "nim",
		},
	}

	return users, nil
}

func (s *testUserServiceImp) Get(id int) (*user.User, error) {
	return &user.User{
		ID:        2,
		FirstName: "NiM 2",
		LastName:  "nim 2",
	}, nil
}

func (s *testUserServiceImp) New(u user.User) error {
	return nil
}

func (s *testUserServiceImp) Update(id int) (*user.User, error) {
	return &user.User{
		ID:        2,
		FirstName: "NiM 2",
		LastName:  "nim 2",
	}, nil
}

func (s *testUserServiceImp) Delete(id int) error {
	return nil
}

func TestAll(t *testing.T) {
	s := &Router{
		userService: &testUserServiceImp{},
	}
	r := setupRouter(s)
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/users/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetByID(t *testing.T) {
	s := &Router{
		userService: &testUserServiceImp{},
	}
	r := setupRouter(s)
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/users/2", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestNewUser(t *testing.T) {
	s := &Router{
		userService: &testUserServiceImp{},
	}

	jsonBody := `{"first_name":"NiM","last_name":"nim"}`
	r := setupRouter(s)
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/users/", strings.NewReader(jsonBody))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateUser(t *testing.T) {
	s := &Router{
		userService: &testUserServiceImp{},
	}

	jsonBody := `{"first_name":"NiM","last_name":"nim"}`
	r := setupRouter(s)
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/users/2", strings.NewReader(jsonBody))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteUser(t *testing.T) {
	s := &Router{
		userService: &testUserServiceImp{},
	}
	r := setupRouter(s)
	req, _ := http.NewRequest(http.MethodDelete, "http://localhost:8080/users/2", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAllAccount(t *testing.T) {
	s := &Router{
		accountService: &testAccountServiceImp{},
	}
	r := setupRouter(s)
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/users/2/bankAccounts/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestNewAccount(t *testing.T) {
	s := &Router{
		accountService: &testAccountServiceImp{},
	}

	jsonBody := `{"account_number":"1234567890","name":"Nim"}`
	r := setupRouter(s)
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/users/2/bankAccounts/", strings.NewReader(jsonBody))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteAccount(t *testing.T) {
	s := &Router{
		accountService: &testAccountServiceImp{},
	}
	r := setupRouter(s)
	req, _ := http.NewRequest(http.MethodDelete, "http://localhost:8080/bankAccounts/1234567890", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeposit(t *testing.T) {
	s := &Router{
		accountService: &testAccountServiceImp{},
	}

	jsonBody := `{"amount":"100"}`
	r := setupRouter(s)
	req, _ := http.NewRequest(http.MethodPut, "http://localhost:8080/bankAccounts/1234567890/deposit", strings.NewReader(jsonBody))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestWithdraw(t *testing.T) {
	s := &Router{
		accountService: &testAccountServiceImp{},
	}

	jsonBody := `{"amount":"100"}`
	r := setupRouter(s)
	req, _ := http.NewRequest(http.MethodPut, "http://localhost:8080/bankAccounts/1234567890/withdraw", strings.NewReader(jsonBody))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
