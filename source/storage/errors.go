package storage

import "errors"

const (
	errDuplicate = "duplicate"
	errNotFound  = "not found"
)

// ErrDuplicate is thrown when a uniqueness constraint is violated.
var ErrDuplicate = errors.New(errDuplicate)

// ErrNotFound is thrown when the searched for model (or object) could not be
// found. Typically used when e.g. get by id operations are performed.
var ErrNotFound = errors.New(errNotFound)

// IsErrDuplicate checks whether an error is of type IsErrDuplicate.
func IsErrDuplicate(err error) bool {
	return err.Error() == errDuplicate
}

// IsErrNotFound checks whether an error is of type ErrNotFound.
func IsErrNotFound(err error) bool {
	return err.Error() == errNotFound
}
