package account

import (
	"os"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Account struct {
	ID            int    `bson:"id" json:"id"`
	UserID        int    `bson:"user_id" json:"user_id"`
	AccountNumber string `bson:"account_number" json:"account_number" binding:"required"`
	Name          string `bson:"name" json:"name" binding:"required"`
	Balance       int    `bson:"balance" json:"balance"`
}

type AccountService interface {
	All(user int) ([]Account, error)
	New(user int, account Account) error
	Deposit(id int, amount int) (*Account, error)
	Withdraw(id int, amount int) (*Account, error)
	Delete(id int) error
	Transfer(fromID int, toID int, amount int) (*Account, error)
}

type AccountServiceImp struct {
	Session *mgo.Session
}

func (s *AccountServiceImp) All(user int) ([]Account, error) {
	sess, err := mgo.Dial(os.Getenv("MLAB_URI"))
	if err != nil {
		return nil, err
	}
	defer sess.Close()

	c := sess.DB("go-training-account").C("Account")
	var results []Account
	err = c.Find(bson.M{"user_id": user}).All(&results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (s *AccountServiceImp) New(user int, account Account) error {
	return nil
}

func (s *AccountServiceImp) Deposit(id int, amount int) (*Account, error) {
	return &Account{}, nil
}

func (s *AccountServiceImp) Withdraw(id int, amount int) (*Account, error) {
	return &Account{}, nil
}

func (s *AccountServiceImp) Delete(id int) error {
	return nil
}

func (s *AccountServiceImp) Transfer(fromID int, toID int, amount int) (*Account, error) {
	return &Account{}, nil
}
