package user

import (
	"os"

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
}

// mongodb://testuser:user1234@ds253203.mlab.com:53203/go-training-account

func (s *UserServiceImp) All() ([]User, error) {
	sess, err := mgo.Dial(os.Getenv("MLAB_URI"))
	if err != nil {
		return nil, err
	}
	defer sess.Close()

	c := sess.DB("go-training-account").C("User")

	if err != nil {
		return nil, err
	}

	var results []User
	err = c.Find(nil).All(&results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (s *UserServiceImp) Get(id int) (*User, error) {
	sess, err := mgo.Dial(os.Getenv("MLAB_URI"))
	if err != nil {
		return nil, err
	}
	defer sess.Close()

	c := sess.DB("go-training-account").C("User")

	if err != nil {
		return nil, err
	}

	var result User
	err = c.Find(bson.M{"id": id}).One(&result)

	return &result, err
}

func (s *UserServiceImp) New(u User) error {
	return nil
}

func (s *UserServiceImp) Update(u User) (*User, error) {
	return &User{}, nil
}

func (s *UserServiceImp) Delete(id int) error {
	return nil
}
