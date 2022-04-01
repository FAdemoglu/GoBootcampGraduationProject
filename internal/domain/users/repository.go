package users

import "gorm.io/gorm"

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Migration() error {
	return r.db.AutoMigrate(&User{})
}

func (r *UserRepository) CreateUser(username, password string) error {
	req := User{
		Username: username,
		Password: password,
		Roles:    "customer",
	}
	result := r.db.Where(User{Username: req.Username}).Attrs(User{Username: req.Username, Password: req.Password, Roles: req.Roles}).FirstOrCreate(&req)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepository) GetByUsername(username string) *User {
	var users *User
	r.db.Where("username = ?", username).Find(&users)
	return users
}

func (r *UserRepository) GetByUsernameAndPassword(username, password string) *User {
	var users *User
	r.db.Where("username = ? AND password = ?", username, password).Find(&users)
	return users
}

func (r *UserRepository) InsertSampleData() {
	users := []*User{
		{
			Id:       1,
			Username: "Furkan Ademoglu",
			Password: "1234",
			Roles:    "admin",
		},
		{
			Id:       2,
			Username: "customer",
			Password: "12345",
			Roles:    "customer",
		},
	}
	for _, u := range users {
		r.db.Where(User{Username: u.Username}).Attrs(User{Username: u.Username, Id: u.Id, Password: u.Password, Roles: u.Roles}).FirstOrCreate(&u)
	}
}
