package database_handler

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

func NewMySQLDB(connection string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{
		PrepareStmt: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NoLowerCase:   true,
		},
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})

	if err != nil {
		panic(fmt.Sprintf("Cannot connect to database :%s", err.Error()))
	}
	return db
}
