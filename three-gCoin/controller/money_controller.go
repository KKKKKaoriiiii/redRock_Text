package controller

import (
	"github.com/gin-gonic/gin"
	"message-broad/service"
	"message-broad/tool"
	"strconv"
)

// MoneyRouter 注册路由
func MoneyRouter(r *gin.Engine) {
	r.POST("/transfer", check, transfer)
	r.POST("/invest", check, invest)
	r.POST("findRecord", check, findRecord)
}

// transfer 转账
func transfer(c *gin.Context) {
	which := c.PostForm("toWho")
	money := c.PostForm("money")
	remarks := c.DefaultPostForm("remarks", "")
	username := tool.CheckLog(c)
	db := tool.GetDb()
	myMoney, err := service.FindMyMoney(db, username)
	tool.CheckErr(err)
	if err != nil {
		return
	}
	moneyNum, _ := strconv.Atoi(money)
	if myMoney < moneyNum {
		tool.PrintInfo(c, "你的余额不足！")
		return
	}
	err = service.ChangeMoney(db, username, which, moneyNum)
	tool.CheckErr(err)
	if err != nil {
		return
	}
	if remarks != "" {
		err = service.AddRecord(db, username, which, moneyNum)
		tool.CheckErr(err)
		if err != nil {
			return
		}
	}
}

// invest 充值
func invest(c *gin.Context) {
	money := c.PostForm("money")
	username := tool.CheckLog(c)
	moneyNum, _ := strconv.Atoi(money)
	db := tool.GetDb()
	err := service.AddMoney(db, username, moneyNum)
	tool.CheckErr(err)
	if err != nil {
		return
	}
}

// findRecord 查找转账记录
func findRecord(c *gin.Context) {
	remarks := c.PostForm("remarks")
	db := tool.GetDb()
	u, err := service.FindRecord(db, remarks)
	tool.CheckErr(err)
	if err != nil {
		return
	}
	tool.PrintRecord(c, u)
}
