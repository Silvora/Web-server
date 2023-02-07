package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//自定义一个字符串
var jwtkey = []byte("757909414@qq.com")

//var str string

type Claims struct {
	UserId string
	jwt.StandardClaims
}

//颁发token
func SetToken(id string) (string, error) {
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "Zjs-7579",   // 签名颁发者
			Subject:   "user token", //签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// fmt.Println(token)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		fmt.Println(err)
	}
	// str = tokenString
	// ctx.JSON(200, gin.H{"token": tokenString})

	return tokenString, nil
}

//解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})
	return token, Claims, err
}

//检验token
func GetTokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("ZJS-7579-Admin-Token")
		//Get("Authorization")
		url := c.Request.URL.Path
		//fmt.Println(token)
		// var imgUrl = regexp.MustCompile(`/images`)
		// var videoUrl = regexp.MustCompile(`/videos`)
		// var wsUrl = regexp.MustCompile(`/ws`)

		// if url == "/root/login" || url == "/root/addUser" || imgUrl.MatchString(url) || videoUrl.MatchString(url) || wsUrl.MatchString(url) {
		// 	//c.Abort()
		// 	return
		// }
		if url == "/blog/login" {
			return
		}
		if url == "/blog/tag" {
			return
		}
		if url == "/blog/class" {
			return
		}
		if url == "/blog/markdown" {
			return
		}
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				//"msg": "当前未登录系统, 无权限访问子应用",
				"msg":  "请求未携带token, 无权限访问",
				"code": 401,
			})
			c.Abort()
			return
		}
		// parseToken 解析token包含的信息
		_, claims, err := ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				//"msg": err.Error(),
				"msg":  "身份信息已过期",
				"code": 403,
			})
			c.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
	}
}
