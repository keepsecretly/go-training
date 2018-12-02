package account

import (
	"errors"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Account struct {
	ID            int    `bson:"id" json:"id"`
	UserID        int    `bson:"user_id" json:"user_id" binding:"required`
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
	Transfer(fromID string, toID string, amount int) (*Account, error)
}

type AccountServiceImp struct {
	Session *mgo.Session
}

func (s *AccountServiceImp) collection() *mgo.Collection {
	return s.Session.DB("go-training-account").C("Account")
}

func (s *AccountServiceImp) All(user int) ([]Account, error) {
	var results []Account
	err := s.collection().Find(bson.M{"user_id": user}).All(&results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (s *AccountServiceImp) New(user int, account Account) error {
	var result Account

	err := s.collection().Find(bson.M{"account_number": account.AccountNumber}).One(&result)

	if err == nil {
		return errors.New("Duplicated account")
	}

	if err != nil && err != mgo.ErrNotFound {
		return err
	}

	err = s.collection().Find(nil).Sort("-id").One(&result)

	if err != nil {
		return err
	}

	account.UserID = user
	account.ID = result.ID + 1

	return s.collection().Insert(account)
}

func (s *AccountServiceImp) Deposit(id int, amount int) (*Account, error) {
	var result Account
	err := s.collection().Find(bson.M{"id": id}).One(&result)

	if err != nil {
		return nil, err
	}

	result.Balance += amount
	err = s.collection().Update(bson.M{"id": id}, result)
	return &result, err
}

func (s *AccountServiceImp) Withdraw(id int, amount int) (*Account, error) {
	var result Account
	err := s.collection().Find(bson.M{"id": id}).One(&result)

	if err != nil {
		return nil, err
	}

	if amount < result.Balance {
		result.Balance -= amount
	} else {
		result.Balance = 0
	}

	err = s.collection().Update(bson.M{"id": id}, result)
	return &result, err
}

func (s *AccountServiceImp) Delete(id int) error {
	err := s.collection().Remove(bson.M{"id": id})
	return err
}

func (s *AccountServiceImp) Transfer(fromID string, toID string, amount int) (*Account, error) {
	var fromAccount Account
	var toAccount Account

	err := s.collection().Find(bson.M{"account_number": fromID}).One(&fromAccount)

	if err != nil {
		return nil, err
	}

	err = s.collection().Find(bson.M{"account_number": toID}).One(&toAccount)

	if err != nil {
		return nil, err
	}

	if amount <= 0 {
		return nil, errors.New("Invalid amount")
	}

	_, err = s.Withdraw(fromAccount.ID, amount)
	if err != nil {
		return nil, err
	}

	_, err = s.Deposit(toAccount.ID, amount)
	if err != nil {
		return nil, err
	}

	if err := s.collection().Find(bson.M{"account_number": fromID}).One(&fromAccount); err != nil {
		return nil, err
	}

	log.Printf("%#v", fromAccount)

	return &fromAccount, err
}
