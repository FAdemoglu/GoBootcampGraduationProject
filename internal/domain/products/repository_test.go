package products

import (
	"github.com/FAdemoglu/graduationproject/pkg/database_handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"testing"
)

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

type Suite struct {
	suite.Suite
	db         *gorm.DB
	repository *ProductRepository
	city       *Product
}

func (s *Suite) SetupSuite() {
	conString := "root:Furkan7937.@tcp(127.0.0.1:3306)/graduation_test?parseTime=True&loc=Local"
	s.db = database_handler.NewMySQLDB(conString)
	s.repository = NewProductRepository(s.db)
	// Migrate Table
	for _, val := range getModels() {
		s.db.AutoMigrate(val)
	}
}

func (s *Suite) TestCityRepository_CreateCity() {
	tests := []struct {
		tag     string
		product *Product
	}{
		{"Ipad", NewProduct("Ipad", 100, 10, 2, 111118)},
	}
	for _, test := range tests {
		err := s.repository.CreateProduct(*test.product)
		assert.Equal(s.T(), nil, err, "Error should be nil")
	}

}

func (s *Suite) TestCityRepository_DeleteCityById() {
	tests := []struct {
		tag string
		Id  int
	}{
		{"With 1 ", 1},
		{"With 10", 10},
	}
	for _, test := range tests {
		err := s.repository.DeleteProductById(test.Id)
		assert.Equal(s.T(), nil, err, "Error should be nil")
	}

}

func (t *Suite) TearDownSuite() {
	sqlDB, _ := t.db.DB()
	defer sqlDB.Close()

	// Drop Table
	for _, val := range getModels() {
		t.db.Migrator().DropTable(val)
	}
}

func getModels() []interface{} {
	return []interface{}{&Product{}}
}
