package users

type User struct {
	Id       int `gorm:"column:UserId,primary_key;auto_increment;not_null"`
	Username string
	Password string
	Roles    string
}

func (User) TableName() string {
	return "User"
}
