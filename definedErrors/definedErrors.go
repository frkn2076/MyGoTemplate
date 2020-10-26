package definedErrors

import(
	"errors"
)

var UserNotFound error
var UserAlreadyExists error
var NotAValidEmail error

func init() {
	UserNotFound = errors.New("Err-1")
	UserAlreadyExists = errors.New("Err-2")
	NotAValidEmail = errors.New("Err-3")
}
