package service

import (
	"database/sql"
	"message-broad/Struct"
	"message-broad/dao"
	"message-broad/tool"
)

func FindMyMoney(db *sql.DB, username string) (int, error) {
	money, err := dao.FindMyMoney(db, username)
	return money, err
}

func ChangeMoney(db *sql.DB, username string, toWho string, money int) error {
	err := dao.ChangeMoney(db, username, toWho, money)
	tool.CheckErr(err)
	return err
}

func AddRecord(db *sql.DB, username string, toWho string, money int) error {
	err := dao.AddRecord(db, username, toWho, money)
	tool.CheckErr(err)
	return err
}

func AddMoney(db *sql.DB, username string, money int) error {
	MyMoney, err := dao.FindMyMoney(db, username)
	tool.CheckErr(err)
	if err != nil {
		return err
	}
	MyMoney += money
	err = dao.GetMoney(db, username, MyMoney)
	tool.CheckErr(err)
	if err != nil {
		return err
	}
	return nil
}

func FindRecord(db *sql.DB, remarks string) (Struct.Record, error) {
	record, err := dao.FindRecord(db, remarks)
	return record, err
}
