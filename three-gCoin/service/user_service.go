package service

import (
	"database/sql"
	"fmt"
	"message-broad/Struct"
	"message-broad/dao"
)

// QueryRowDemoPassword 检验密码和账号是否正确。
func QueryRowDemoPassword(db *sql.DB, username string, pwd string) bool {
	sqlStr := "select username, password from user where username = ?"
	row := db.QueryRow(sqlStr, username)
	err := row.Err()
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	} else {
		var usr, password string
		err = row.Scan(&usr, &password)
		if err != nil {
			fmt.Println(err)
		}
		if usr == username && pwd == password {
			return true
		} else {
			return false
		}
	}
	return false
}

// RegisterUser 注册服务
func RegisterUser(db *sql.DB, username string, password string, answer string, question string) bool {
	err := dao.InsertUser(db, username, password, answer, question)
	if err != nil {
		fmt.Println("Insert failed err:", err)
		return false
	}
	return true
}

// QueryRowDemo 查找用户
func QueryRowDemo(db *sql.DB, username string) (Struct.User, bool) {
	u, flag := dao.FindUser(db, username)
	return u, flag
}

// AddSecret 添加密保
func AddSecret(db *sql.DB, question string, answer string, username string) error {
	err := dao.AddSecret(db, question, answer, username)
	return err
}
