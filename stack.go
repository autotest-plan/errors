package errors

import "fmt"

type GeneralError struct {
	errorCode int
	message   string
}

func (ge *GeneralError) Error() string {
	return fmt.Sprintf("error(%d):%s\n", ge.errorCode, ge.message)
}

func Error(errCode int, message string) error {
	return &GeneralError{
		errorCode: errCode,
		message:   message,
	}
}
