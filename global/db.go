package global

import (
	"Blob/config"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

func NewDB(db_ *config.Database) (*gorm.DB, error) {
	db, err := gorm.Open(db_.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		db_.Username,
		db_.Password,
		db_.Host,
		db_.Project))

	if err != nil {
		return nil, errors.Wrap(err, "Init a new db error")
	}

	if Server.Mode == "Debug" {
		db.LogMode(true)
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(db_.MaxIdleConns)
	db.DB().SetMaxOpenConns(db_.MaxOpenConns)

	return db, nil
}
