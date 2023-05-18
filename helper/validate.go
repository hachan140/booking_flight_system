package helper

import (
	"errors"
	"regexp"
)

func StringValidation(input string, format string, msg string) error {
	if m, _ := regexp.MatchString(format, input); !m {
		return errors.New(msg + " does not match")
	}
	return nil
}
