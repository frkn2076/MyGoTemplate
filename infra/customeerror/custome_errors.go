package customeerror

import (
	"errors"
)

var UserNotFound error
var UserAlreadyExists error
var NotAValidEmail error
var WrongPassword error
var ShortPassword error

func init() {
	UserNotFound = errors.New("Err-1")
	UserAlreadyExists = errors.New("Err-2")
	NotAValidEmail = errors.New("Err-3")
	WrongPassword = errors.New("Err-4")
	ShortPassword = errors.New("Err-5")
}
