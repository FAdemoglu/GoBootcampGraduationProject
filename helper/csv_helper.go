package helper

import (
	"encoding/csv"
	"github.com/FAdemoglu/graduationproject/internal/domain/category"
	"os"
)

func ReadCsvToBookSlice(fileName string) ([]category.Category, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	reader := csv.NewReader(f)

	lines, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	var result []category.Category

	for _, line := range lines[1:] {
		data := category.Category{
			CategoryName: line[0],
		}
		result = append(result, data)
	}

	return result, nil

}
