package validate

import (
	"errors"

	"github.com/fatih/structs"
)

func ValidateUser(obj interface{}) error {
	values := structs.Values(obj)

	for _, v := range values {
		if v == "" {
			return errors.New("object has a string attribute empty")
		}
	}

	return nil
}
