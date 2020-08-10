package sqlx

import "fmt"

type Error struct {
	err string
}

// New
func New(message string, args ...interface{}) error {
	return Error{err: "goqu: " + fmt.Sprintf(message, args...)}
}

// NewEncodeError
func NewEncodeError(t interface{}) error {
	return Error{err: "goqu_encode_error: " + fmt.Sprintf("Unable to encode value %+v", t)}
}

func (e Error) Error() string {
	return e.err
}
