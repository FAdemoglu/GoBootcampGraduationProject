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

func (r *UserRepository) CreateUser(u *User) error {
	result := r.db.Create(u)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepository) GetByUsername(username string) []User {
	var users []User
	r.db.Where("Name LIKE ?", "%"+username+"%").Find(&users)
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
		r.db.Create(&u)
	}
}
