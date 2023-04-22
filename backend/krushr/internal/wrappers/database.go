package wrappers

import (
	"errors"

	"gorm.io/gorm"
)

// WithTransaction starts a transaction with the provided database instance and executes a function that is passed as an argument.
// The function should have a transaction object as its input parameter and should return two values - a result and an error.
// If the provided database instance has an error, the function returns an empty result and the error that was encountered.
// If the connection pool is unavailable, an error message is returned.
// The function starts a new transaction using the Begin() method of the provided database instance.
// If this method returns an error, an empty result and the error encountered is returned.
// The passed function is executed inside a deferred function.
// If the function panics, the transaction is rolled back and the panic is propagated.
// If the passed function returns an error, the transaction is rolled back, and the error is returned.
// If there is a panic in the deferred function, the transaction is rolled back, and the panic is propagated.
// If the transaction object has an error, the transaction is rolled back, and the error is returned.
// Otherwise, the transaction is committed and the result of the function execution is returned.
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
