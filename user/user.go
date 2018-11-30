package user

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `bson:"first_name" json:"first_name" binding:"required"`
	LastName  string `bson:"last_name" json:"last_name" binding:"required"`
}

type UserService interface {
	All() ([]User, error)
	Get(id int) (*User, error)
	New(u User) error
	Update(u User) (*User, error)
	Delete(id int) error
}

type UserServiceImp struct {
	Session *mgo.Session
}

// mongodb://testuser:user1234@ds253203.mlab.com:53203/go-training-account

func (s *UserServiceImp) All() ([]User, error) {
	c := s.Session.DB("go-training-account").C("User")
	var results []User
	err := c.Find(nil).All(&results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (s *UserServiceImp) Get(id int) (*User, error) {
	c := s.Session.DB("go-training-account").C("User")
	var result User
	err := c.Find(bson.M{"id": id}).One(&result)

	return &result, err
}

func (s *UserServiceImp) New(u User) error {
	c := s.Session.DB("go-training-account").C("User")
	var result User
	err := c.Find(nil).Sort("-id").One(&result)

	if err != nil {
		return err
	}

	u.ID = result.ID + 1

	return c.Insert(u)
}

func (s *UserServiceImp) Update(u User) (*User, error) {
	c := s.Session.DB("go-training-account").C("User")
	err := c.Update(bson.M{"id": u.ID}, u)

	return &u, err
}

func (s *UserServiceImp) Delete(id int) error {
	c := s.Session.DB("go-training-account").C("User")
	err := c.Remove(bson.M{"id": id})

	return err
}
