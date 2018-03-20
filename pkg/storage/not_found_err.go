package storage

import (
	"fmt"
)

type ErrNotFound struct {
	BasePath string
	Module   string
}

func (n ErrNotFound) Error() string {
	return fmt.Sprintf("%s/%s not found", n.BasePath, n.Module)
}

type ErrVersionNotFound struct {
	BasePath string
	Module   string
	Version  string
}

func (e ErrVersionNotFound) Error() string {
	return fmt.Sprintf("%s/%s@%s not found", e.BasePath, e.Module, e.Version)
}

func IsNotFoundError(err error) bool {
	_, ok := err.(ErrNotFound)
	return ok
}

type ErrVersionAlreadyExists struct {
	BasePath string
	Module   string
	Version  string
}

func (e ErrVersionAlreadyExists) Error() string {
	return fmt.Sprintf("%s/%s@%s", e.BasePath, e.Module, e.Version)
}

func IsVersionAlreadyExistsErr(err error) bool {
	_, ok := err.(ErrVersionAlreadyExists)
	return ok
}
