package users

type UserService struct {
	u UserRepository
}

func NewUserService(u UserRepository) *UserService {
	return &UserService{
		u: u,
	}
}
