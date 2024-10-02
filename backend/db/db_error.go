package db

import "errors"

type DBError struct {
	Err error
}

func (e *DBError) Error() string { return e.Err.Error() }

func NewDBError(msg string) error { return &DBError{Err: errors.New(msg)} }
