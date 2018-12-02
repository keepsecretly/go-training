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

func (s *UserServiceImp) collection() *mgo.Collection {
	return s.Session.DB("go-training-account").C("User")
}

func (s *UserServiceImp) All() ([]User, error) {
	var results []User
	err := s.collection().Find(nil).All(&results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (s *UserServiceImp) Get(id int) (*User, error) {
	var result User
	err := s.collection().Find(bson.M{"id": id}).One(&result)
	return &result, err
}

func (s *UserServiceImp) New(u User) error {
	var result User
	err := s.collection().Find(nil).Sort("-id").One(&result)

	if err != nil {
		return err
	}

	u.ID = result.ID + 1
	return s.collection().Insert(u)
}

func (s *UserServiceImp) Update(u User) (*User, error) {
	err := s.collection().Update(bson.M{"id": u.ID}, u)
	return &u, err
}

func (s *UserServiceImp) Delete(id int) error {
	err := s.collection().Remove(bson.M{"id": id})
	return err
}
