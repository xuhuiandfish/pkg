package db

import (
	"errors"

	gorm "gorm.io/gorm"
)

func IsErrorNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

// ErrOptimisticLock is returned by if the struct being
// modified has a Version field and the value is not equal
// to the current value in the database
var ErrOptimisticLock = errors.New("optimistic Lock Error")

var ErrRecordNotFound = gorm.ErrRecordNotFound
