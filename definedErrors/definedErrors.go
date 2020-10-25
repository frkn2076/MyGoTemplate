package definedErrors

import(
	"errors"
)

var UserNotFound error
var UserAlreadyExists error

func init() {
	UserNotFound = errors.New("Err-1")
	UserAlreadyExists = errors.New("Err-2")
}
