package common

import "fmt"

const (
	AwsProvider = iota
	GoogleProvider
	DigitalOcean
)

type Error struct {
	Code int
	Description string
	Function string
	Package string
}

const (
	InvalidProviderErrCode = iota + 1
)

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s, function: %s, code: %d", e.Package, e.Description, e.Function, e.Code)
}

func NewError(code int, fn, description, pkg string) *Error {
	return &Error{code, fn, description, pkg}
}

func InvalidProviderErr(fn, pkg string) *Error {
	return NewError(InvalidProviderErrCode, fn, "No valid error was provided", pkg)
}