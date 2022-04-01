package users

type User struct {
	Id       int `gorm:"column:UserId;autoIncrement;PRIMARY_KEY;not null"`
	Username string
	Password string
	Roles    string
}

func (User) TableName() string {
	return "User"
}
