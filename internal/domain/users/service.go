package users

type UserService struct {
	u UserRepository
}

func NewUserService(u UserRepository) *UserService {
	return &UserService{
		u: u,
	}
}

func (u *UserService) Create(user *User) error {
	existUser := u.u.GetByUsername(user.Username)
	if existUser != nil && len(existUser) > 0 {
		return ErrUserExistWithUsername
	}
	err := u.u.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}
