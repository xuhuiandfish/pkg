package db

import (
	"github.com/hashicorp/go-multierror"
	gorm "gorm.io/gorm"
)

type DB struct {
	write *gorm.DB
	read  *gorm.DB
}

func (db *DB) Update() *gorm.DB {
	return db.write
}

func (db *DB) View() *gorm.DB {
	return db.read
}

func (db *DB) Debug() *DB {
	return &DB{
		write: db.write.Debug(),
		read:  db.read.Debug(),
	}
}

func (db *DB) Begin() *DB {
	tx := db.write.Begin()

	return &DB{
		write: tx,
		read:  db.read,
	}
}

func (db *DB) Rollback() error {
	return db.write.Rollback().Error
}

func (db *DB) Commit() error {
	return db.write.Commit().Error
}

// Deprecated in grom 2.0
// func (db *DB) RollbackUnlessCommitted() {
// 	if err := db.write.RollbackUnlessCommitted().Error; err != nil {
// 		logrus.WithError(err).Error("DB: RollbackUnlessCommitted")
// 	}
// }

func (db *DB) Tx(fn func(tx *gorm.DB) error) error {
	return db.write.Transaction(fn)
}

func (db *DB) Ping() error {
	database, err := db.write.DB()
	if err != nil {
		return err
	}
	return database.Ping()
}

func (db *DB) Close() error {
	var merr *multierror.Error
	if write, err := db.write.DB(); err != nil {
		merr = multierror.Append(merr, err)
	} else {
		if err := write.Close(); err != nil {
			merr = multierror.Append(merr, err)
		}
	}

	if read, err := db.read.DB(); err != nil {
		merr = multierror.Append(merr, err)
	} else {
		if err := read.Close(); err != nil {
			merr = multierror.Append(merr, err)
		}
	}

	return merr.ErrorOrNil()
}
