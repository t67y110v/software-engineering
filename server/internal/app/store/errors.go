package store

import "errors"

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEmptyValue     = errors.New("empty colums in select")
)
