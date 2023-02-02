package service

import (
	"indexServer/logger"
	"indexServer/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserEmail struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

func UserToEmail(ctx *gin.Context) {

	var userEmail UserEmail

	if err := ctx.ShouldBindJSON(&userEmail); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		logger.SetLogger(1, "解析json数据出错")

		ctx.JSON(http.StatusOK, gin.H{
			"code":  200,
			"error": err.Error(),
		})
		return
	}

	isEmail := middleware.UserToMyEmail(userEmail.Name, userEmail.Email, userEmail.Message)

	if isEmail {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"isEmail": true,
			"message": "邮箱发送成功",
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"isEmail": false,
			"message": "邮箱发送失败",
		})
		return
	}

}
