package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-broad/dao"
	"message-broad/service"
	"message-broad/tool"
	"net/http"
)

// UserRouter 注册路由
func UserRouter(r *gin.Engine) {
	r.POST("/login", login)
	r.POST("/register", register)
	r.POST("/change", changePassword)
	r.GET("/logout", logout)
	r.POST("/showSecret", showSecret)
	r.POST("/addSecret", addSecret)
}

// check 检查登录中间件
func check(c *gin.Context){
	username := tool.CheckLog(c)
	if username == "" {
		tool.PrintInfo(c, "你还没有登录!")
		return
	}
	c.Next()
}

// addSecret 添加密保
func addSecret(c *gin.Context) {
	username := tool.CheckLog(c)
	db := tool.GetDb()
	u, _ := dao.FindUser(db, username)
	if u.Question == "" || u.Answer == "" {
		question := c.PostForm("question")
		answer := c.PostForm("answer")
		err:=service.AddSecret(db, question, answer, username)
		tool.CheckErr(err)
		return
	}
	tool.PrintInfo(c, "你已经设置密保！")
}

// login 登录
func login(c *gin.Context) {
	username := tool.CheckLog(c)
	if username != "" {
		c.String(200, "你已经登录，请先退出登录！")
		return
	}
	username = c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println("LoginUserInfo:", username, password)
	db := tool.GetDb()
	flag := service.QueryRowDemoPassword(db, username, password)
	if flag {
		c.SetCookie("login", username, 3600, "/", "", false, true)
		tool.PrintInfo(c, "登陆成功!")
		return
	}
	tool.PrintInfo(c, "账号或密码错误！")
}

// register 注册
func register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	secretQuestion := c.DefaultPostForm("question", "")
	secretAnswer := c.DefaultPostForm("answer", "")
	db := tool.GetDb()
	_, flag := service.QueryRowDemo(db, username)
	if flag {
		tool.PrintInfo(c, "该账号已经被注册!")
		return
	}
	flag2 := service.RegisterUser(db, username, password, secretAnswer, secretQuestion)
	if !flag2 {
		tool.PrintInfo(c, "注册失败！")
		return
	}
	tool.PrintInfo(c, "注册成功！")
}

// changePassword 修改密码
func changePassword(c *gin.Context) {
	username := c.PostForm("username")
	db := tool.GetDb()
	u, flag := service.QueryRowDemo(db, username)
	if flag {
		answer := c.PostForm("answer")
		password := c.PostForm("newPassword")
		if answer == u.Answer {
			err := dao.UpdateRowDemo(db, password, username)
			tool.CheckErr(err)
			tool.PrintInfo(c, "修改完成！")
		}
	} else {
		password2 := c.PostForm("newPassword")
		err := dao.UpdateRowDemo(db, password2, username)
		tool.CheckErr(err)
		tool.PrintInfo(c, "你还没有设置密保，修改完成!")
	}
}

// logout 退出登录
func logout(c *gin.Context) {
	value := tool.CheckLog(c)
	if value == "" {
		tool.PrintInfo(c, "未登录 ")
		return
	}
	cookie, err := c.Request.Cookie("login")
	tool.CheckErr(err)
	cookie.MaxAge = -1
	http.SetCookie(c.Writer, cookie)
	tool.PrintInfo(c, "退出登录成功")
	return
}

// showSecret 展示密保
func showSecret(c *gin.Context) {
	username := c.PostForm("username")
	db := tool.GetDb()
	u, flag := service.QueryRowDemo(db, username)
	if !flag {
		tool.PrintInfo(c, "查无此号")
		return
	}
	tool.PrintInfo(c, "你的密保问题为："+u.Question)
}
