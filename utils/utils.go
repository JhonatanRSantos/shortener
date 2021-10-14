package utils

import (
	"errors"
	"math/rand"
	"time"

	"github.com/go-playground/validator/v10"
)

const timeZoneLocation = "America/Sao_Paulo"

// Parse and validate body params
func ParseAndValidateBodyParams(bodyParser func(out interface{}) error, params interface{}) error {
	if bodyParser(params) != nil {
		return errors.New("failed to parse request")
	} else if validator.New().Struct(params) != nil {
		return errors.New("the request has missing params")
	}
	return nil
}

// Get the current date
func GetCurrenteDate() (time.Time, error) {
	location, err := time.LoadLocation(timeZoneLocation)

	if err != nil {
		return time.Time{}, errors.New("failed to proccess the request date")
	}
	return time.Now().In(location), nil
}

// Generates a random int between min and max
func RandomInt(min int, max int) int {
	time.Sleep(time.Nanosecond * 2)
	return min + rand.New(
		rand.NewSource(time.Now().UnixNano()),
	).Intn(max-min+1)
}

// Generates a random String
func RandomString(size int) string {
	result := []rune{}

	for {
		result = append(result, RandomRune())
		if len(result) == size {
			return string(result)
		}
	}
}

// Generates a random rune
func RandomRune() rune {
	switch RandomInt(1, 3) {
	case 1:
		return rune(RandomInt(48, 57))
	case 2:
		return rune(RandomInt(65, 90))
	default:
		return rune(RandomInt(97, 122))
	}
}
