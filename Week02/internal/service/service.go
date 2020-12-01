package service

import (
	"database/sql"
	"errors"
	"handleError/internal/dao"
)


func GetUserInfo(name string) (r dao.UserInfo, err error) {
	statement := "select * from info_user where name = ?"
	r, err = dao.SearchOneRecord(statement, name)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	//处理ErrNoRows，将error设置为nil, 给数据添加默认值
	if errors.Is(err, sql.ErrNoRows) {
		err = nil
		r.UserId = "xxx"
		r.Name = "吕布"
		r.Age = 24
		r.Comments = "吕布奉先"
		r.Comments = "吾乃三国武将"
	}
	return
}
