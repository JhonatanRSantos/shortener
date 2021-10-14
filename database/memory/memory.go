package memory

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/JhonatanRSantos/shortener/database/interfaces"
)

type MemoryDB struct {
	Database map[string]interface{}
}

// Save
func (mDB *MemoryDB) SaveData(tableName string, data interface{}) error {
	primaryKey, err := mDB.getPrimaryKey(tableName, data)

	if err != nil {
		return err
	}

	(*mDB).Database[primaryKey] = data
	log.Println("Save = ", (*mDB).Database[primaryKey])
	return nil
}

// Find
func (mDB *MemoryDB) FindData(tableName string, data interface{}) (interface{}, error) {
	primaryKey, err := mDB.getPrimaryKey(tableName, data)
	log.Println("Find = ", primaryKey, (*mDB).Database[primaryKey])

	if err != nil {
		return nil, err
	}

	if value, ok := (*mDB).Database[primaryKey]; ok {
		return value, nil
	}
	return nil, errors.New("data not found")
}

// Get the primary key
func (mDB *MemoryDB) getPrimaryKey(tableName string, data interface{}) (string, error) {
	primaryKey := tableName
	reflectedData := reflect.ValueOf(data).Elem()
	numberOfFields := reflectedData.NumField()

	for index := 0; index < numberOfFields; index++ {
		fieldTag := reflectedData.Type().Field(index).Tag.Get(interfaces.TagName)
		fieldValue := reflectedData.Field(index)
		// fieldName := reflectedData.Type().Field(index).Name

		if strings.Contains(string(fieldTag), "primary_key") {
			primaryKey += fmt.Sprintf("#%v", fieldValue.Interface())
			break
		}
	}

	if primaryKey == tableName {
		return "", errors.New("missing primary key")
	}
	return primaryKey, nil
}

// Create a new memory DB
func NewMemoryDB() MemoryDB {
	return MemoryDB{
		Database: make(map[string]interface{}),
	}
}
