package repo

import "errors"

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrAlreadyExists  = errors.New("record already exists")
	ErrNoAffected     = errors.New("no records or relations affected")
)
