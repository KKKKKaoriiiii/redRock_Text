package tool

import (
	"github.com/gin-gonic/gin"
	"message-broad/Struct"
)

// PrintInfo 输出字符串
func PrintInfo(c *gin.Context, str string) {
	c.JSON(200, str+"\n")
}

func PrintRecord(c *gin.Context, u Struct.Record) {
	c.JSON(200, gin.H{
		"toWho":   u.ToWhich,
		"fromWho": u.FromWhich,
		"money":   u.Money,
	})
}
