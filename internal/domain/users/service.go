package users

type UserService struct {
	r UserRepository
}

func NewUserService(u UserRepository) *UserService {
	return &UserService{
		r: u,
	}
}

func (u *UserService) GetByUserNameAndPassword(username, password string) *User {
	user := u.r.GetByUsernameAndPassword(username, password)
	return user
}

func (u *UserService) GetByUsername(username string) *User {
	var user *User
	user = u.r.GetByUsername(username)
	return user
}

func (u *UserService) CreateUser(username, password string) error {
	err := u.r.CreateUser(username, password)
	return err
}
