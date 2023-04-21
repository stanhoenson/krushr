package wrappers

import (
	"errors"

	"gorm.io/gorm"
)

func WithTransaction[T any](db *gorm.DB, fn func(tx *gorm.DB) (T, error)) (T, error) {
	var initialT T
	var err error

	if db.Error != nil {
		return initialT, db.Error
	}

	if db.Statement.ConnPool == nil {
		return initialT, errors.New("cannot start transaction: no available connection pool")
	}

	tx := db.Begin()
	if tx.Error != nil {
		return initialT, tx.Error
	}

	initialT, err = fn(tx)
	defer func() {

		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		} else if err != nil {
			tx.Rollback()
		} else if err := recover(); err != nil {
			tx.Rollback()
			panic(err)
		} else if tx.Error != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	return initialT, err

}
