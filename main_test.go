package main

import (
	"go-training/user"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testUserServiceImp struct {
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

func (s *testUserServiceImp) New() (*user.User, error) {
	return &user.User{
		ID:        2,
		FirstName: "NiM 2",
		LastName:  "nim 2",
	}, nil
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
