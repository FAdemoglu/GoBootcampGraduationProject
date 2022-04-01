package users

type User struct {
	Id       int `gorm:"column:UserId"`
	Username string
	Password string
	Roles    string
}

func (User) TableName() string {
	return "User"
}
