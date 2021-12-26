package dao

import (
	"database/sql"
	"message-broad/Struct"
	"message-broad/tool"
)

func FindMyMoney(db *sql.DB, username string) (int, error) {
	var money int
	sqlStr := "select money from moneyInfo where username = ?"
	err := db.QueryRow(sqlStr, username).Scan(&money)
	if err != nil {
		return 0, err
	}
	return money, nil
}

func ChangeMoney(db *sql.DB, username string, toWho string, money int) error {
	MyMoney, err := FindMyMoney(db, username)
	tool.CheckErr(err)
	if err != nil {
		return err
	}
	WMoney, err := FindMyMoney(db, toWho)
	tool.CheckErr(err)
	if err != nil {
		return err
	}
	MyMoney -= money
	WMoney += money
	err = GetMoney(db, username, MyMoney)
	if err != nil {
		return err
	}
	err = GetMoney(db, toWho, WMoney)
	if err != nil {
		return err
	}
	return nil
}

func GetMoney(db *sql.DB, username string, money int) error {
	sqlStr := "update moneyinfo set money=? where username = ?"
	_, err := db.Exec(sqlStr, money, username)
	tool.CheckErr(err)
	return err
}

func AddRecord(db *sql.DB, username string, toWho string, money int) error {
	sqlStr := "insert into record(towhich, fromwhich, money) values (?, ?, ?)"
	_, err := db.Exec(sqlStr, toWho, username, money)
	tool.CheckErr(err)
	return err
}

func FindRecord(db *sql.DB, remarks string) (Struct.Record, error) {
	sqlStr := "select towhich, fromwhich, money from record where remarks = ?"
	var u Struct.Record
	err := db.QueryRow(sqlStr, remarks).Scan(&u.ToWhich, &u.FromWhich, &u.Money)
	tool.CheckErr(err)
	if err != nil {
		return u, err
	}
	return u, nil
}
