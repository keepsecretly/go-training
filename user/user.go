package user

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
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

func (s *UserServiceImp) All() ([]User, error) {
	users := []User{}
	return users, nil
}

func (s *UserServiceImp) Get(id int) (*User, error) {
	return &User{}, nil
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
