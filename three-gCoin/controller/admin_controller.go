package controller

import (
	"github.com/gin-gonic/gin"
	"message-broad/tool"
)

// AdminRouter 注册路由
func AdminRouter(engine *gin.Engine) {
	engine.POST("/userDelete", AdminMiddleWare, userDelete)
	engine.GET("/help", help)
}

// AdminMiddleWare 管理员权限确认
func AdminMiddleWare(c *gin.Context) {
	username := tool.CheckLog(c)
	if username == "cjw" {
		return
	}
	tool.PrintInfo(c, "你不是管理员！")
	c.Abort()
	return
}

// userDelete 删除用户
func userDelete(c *gin.Context) {
	username := c.PostForm("username")
	db := tool.GetDb()
	sqlStr := "delete from user where username = ?"
	_, err := db.Exec(sqlStr, username)
	tool.CheckErr(err)
	tool.PrintInfo(c, "注销成功！")
}

// help 帮助
func help(c *gin.Context) {
	tool.PrintInfo(c, "/transfer 转账\n/invest 充值\n/findRecord 查询记录\n")
	tool.PrintInfo(c, "/login 登录\n/register 注册\n/change 修改密码\n/logout 登出\n/showSecret 展示密保\n/addSecret 添加密保\n")
}
