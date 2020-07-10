package utils

import (
	"errors"
	"gopkg.in/asaskevich/govalidator.v9"
	"strings"
)

func IsStrIn(input *string, rules string, fieldPath string) (bool, error) {
	if input == nil {
		return true, nil
	}

	split := strings.Split(rules, "|")
	msg := strings.Join(split, ", ")

	if govalidator.IsIn(*input, split...) {
		return true, nil
	}
	return false, errors.New("The " + fieldPath + " field must be one of " + msg)

}
