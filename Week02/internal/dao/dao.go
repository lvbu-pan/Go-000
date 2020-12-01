package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type UserInfo struct {
	UserId   string
	Name     string
	Age      int
	Mobile   string
	Comments string
}

func conn() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:123456@tcp(122.112.161.144:8806)/devops")
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	if err = db.Ping(); err != nil {
		return nil, errors.Wrap(err, "")
	}
	return db, nil
}

func SearchOneRecord(statement string, arg ...interface{}) (row UserInfo, err error) {
	var db *sql.DB
	db, err = conn()
	if err != nil {
		return
	}
	err = db.QueryRow(statement, arg...).Scan(&row.UserId, &row.Name, &row.Age, &row.Mobile, &row.Comments)
	if err != nil {
		return row, errors.Wrap(err, "")
	}
	return
}
