package service

import (
	"indexServer/db"
	"indexServer/logger"
	"indexServer/middleware"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Email    string `json:"email" binding:"required"`
	PassWord string `json:"password" binding:"required"`
	Verify   string `json:"verify"`
}

//获取邮箱验证码
func isEmailVerify(verify string) string {

	return db.GetRedis(verify)
}

//判断用户邮箱是否存在
func isUserEmail(email string) bool {

	var name string
	sql := `select email from User where email=?`

	err := db.MysqlDB.QueryRow(sql, email).Scan(&name)
	if err != nil {
		logger.SetLogger(1, email+": Mysql查询出错 -->"+err.Error())
		log.Println(email + ":Mysql查询出错-->" + err.Error())

		return false
	}

	return true

}

//检查账号是否正确
func checkUser(user User) bool {
	var email string
	var password string
	sql := `select email,password from User where email=?`

	err := db.MysqlDB.QueryRow(sql, user.Email).Scan(&email, &password)
	if err != nil {
		logger.SetLogger(1, email+": Mysql查询出错")
		log.Println(email + ":Mysql查询出错")

		return false
	}

	isOk, _ := middleware.ValidatePassword(password, user.PassWord)

	if isOk {
		return true
	} else {
		return false

	}
}

//添加用户
func addUser(user User) bool {
	sql := `insert into User(email,password) values(?,?)`

	userPass, err := middleware.PasswordToMd5(user.PassWord)
	if err != nil {
		logger.SetLogger(1, user.Email+": 密码转换错误")
		log.Println(user.Email + ":密码转换错误")
		return false
	}

	_, err = db.MysqlDB.Exec(sql, user.Email, userPass)
	if err != nil {
		logger.SetLogger(1, user.Email+": Mysql插入出错")
		log.Println(user.Email + ":Mysql插入出错")
		return false
	}
	return true
}

//验证用户邮箱是否注册
func Status(ctx *gin.Context) {

	email := ctx.DefaultQuery("email", "")

	isUser := isUserEmail(email)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"email":   email,
		"isLogin": isUser,
	})

}

//用户登录注册
func Login(ctx *gin.Context) {
	var user User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		logger.SetLogger(1, "解析json数据出错")
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"error":   err.Error(),
			"isLogin": false,
		})
		return
	}

	isUser := isUserEmail(user.Email)

	//用户登录
	if user.Verify == "" && isUser {
		//判断账号是否正确
		isLogin := checkUser(user)

		token, _ := middleware.SetToken(user.Email)

		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"email":   user.Email,
			"isLogin": isLogin,
			"Token":   token,
		})
		return
	}

	//发送邮箱
	if user.Verify == "" && !isUser {
		isEmail := middleware.SendEmail(user.Email)
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"email":   user.Email,
			"isEmail": isEmail,
			"isLogin": false,
		})
		return
	}

	//添加用户
	if user.Verify != "" && !isUser {
		//判断验证码
		verify := isEmailVerify(user.Email)
		if verify == user.Verify {
			//添加用户
			isAdd := addUser(user)
			token := ""
			if isAdd {
				token, _ = middleware.SetToken(user.Email)
			}
			ctx.JSON(http.StatusOK, gin.H{
				"code":    200,
				"email":   user.Email,
				"isLogin": isAdd,
				"Token":   token,
			})
			return

		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    200,
				"email":   user.Email,
				"msg":     "验证码不正确",
				"isLogin": false,
			})
			return
		}
	}
}
