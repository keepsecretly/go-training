package account

type Account struct {
	ID            int    `json:"id"`
	UserID        int    `json:"user_id"`
	AccountNumber string `json:"account_number"`
	Name          string `json:"name"`
	Balance       int    `json:"balance"`
}

type AccountService interface {
	All(user int) ([]Account, error)
	New(user int) (*Account, error)
	Deposit(id int, amount int) (*Account, error)
	Withdraw(id int, amount int) (*Account, error)
	Delete(id int) error
	Transfer(fromID int, toID int) (*Account, error)
}

type AccountServiceImp struct {
}

func (s *AccountServiceImp) All(user int) ([]Account, error) {
	accs := []Account{}
	return accs, nil
}

func (s *AccountServiceImp) New(user int) (*Account, error) {
	return &Account{}, nil
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

func (s *AccountServiceImp) Transfer(fromID int, toID int) (*Account, error) {
	return &Account{}, nil
}
