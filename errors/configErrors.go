package errors

import (
	"errors"
	"fmt"
)

func ErrUnableToParseConfigFile(fileName string) error {
	errorTxt := fmt.Sprintf("Unable to parse config file: %s. Please check and ensure required available fields are available", fileName)
	return errors.New(errorTxt)
}
