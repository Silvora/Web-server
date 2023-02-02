package middleware

import (
	"indexServer/db"
	"indexServer/logger"
	"log"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"

	"github.com/jordan-wright/email"
)

func setHtmlDoc() (string, string) {

	arr := getRandNumber()
	num := arr[0] + arr[1] + arr[2] + arr[3]

	str := `<div style=" position: relative;
    margin: auto;
    overflow: hidden;
    width: 650px;
    height: 480px;">

        <div style="color: #f6f4f3;
        text-align: center;
        text-transform: uppercase;
        font-family: ' Lato', sans-serif; font-size: 0.7em; letter-spacing: 5px; margin-top: 25%;">
            <div style="display: inline-block;
            padding: 20px;
            width: 100px;
            border-radius: 5px;background: #ef2f3c;">
                <div style="font-family: 'Montserrat', sans-serif;
                color: #183059;
                font-size: 4em;">` + arr[0] + `</div>
            </div>
            <div style="display: inline-block;
            padding: 20px;
            width: 100px;
            border-radius: 5px;background: #f6f4f3;
            color: #183059;">
                <div style="font-family: 'Montserrat', sans-serif;
                color: #183059;
                font-size: 4em;">` + arr[1] + `</div>
            </div>
            <div style="display: inline-block;
            padding: 20px;
            width: 100px;
            border-radius: 5px;background: #276fbf;">
                <div style="font-family: 'Montserrat', sans-serif;
                color: #183059;
                font-size: 4em;">` + arr[2] + `</div>
            </div>
            <div style="display: inline-block;
            padding: 20px;
            width: 100px;
            border-radius: 5px;background: #f0a202;">
                <div style="font-family: 'Montserrat', sans-serif;
                color: #183059;
                font-size: 4em;">` + arr[3] + `</div>
            </div>
        </div>
        <h1 style="font-family: 'Lato', sans-serif;
        text-align: center;
        margin-top: 2em;
        font-size: 1em;
        text-transform: uppercase;
        letter-spacing: 5px;
        color: #df1624;">5 minute effectiveness / 5分钟实效</h1>
    </div>
    <div style="position: fixed;
    bottom: 0;
    right: 0;
    text-transform: uppercase;
    padding: 10px;
    font-family: 'Lato', sans-serif; font-size: 0.7em;">
        <p style="letter-spacing: 3px;
        color: #ef2f3c;">my by <span style="color: rgb(13, 101, 233);">zjs</span> ♡
        </p>
    </div>`
	return str, num
}

func getRandNumber() [4]string {

	var randList [4]string

	rand.Seed(time.Now().UnixNano()) //以当前系统时间作为种子参数

	//  产生随机数
	for i := 0; i < 4; i++ {
		//  限制在100以内
		n := rand.Intn(100)
		var s string
		if n < 10 {
			s = strconv.Itoa(n)
			s = "0" + s
		} else {
			s = strconv.Itoa(n)
		}

		randList[i] = s

	}

	return randList
}

func initEmail(toEmail string, html string, subject string) bool {

	e := email.NewEmail()
	//设置发送方的邮箱
	e.From = "<757909414@qq.com>"
	// 设置接收方的邮箱
	e.To = []string{toEmail}
	//设置主题
	e.Subject = subject
	//设置文件发送的内容
	//e.Text = []byte("www.topgoer.com是个不错的go语言中文文档")
	//str, num := setHtmlDoc()

	e.HTML = []byte(html)

	//设置服务器相关的配置
	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", "757909414@qq.com", "yncowehrgndibejh", "smtp.qq.com"))
	if err != nil {
		logger.SetLogger(1, toEmail+"：邮箱发送失败")
		log.Fatal(err)
		return false
	}

	return true

}

func SendEmail(toEmail string) bool {
	//e := InitEmail()

	html, num := setHtmlDoc()

	bool := initEmail(toEmail, html, "邮箱验证码 Zjs-7579")

	db.SetRedis(toEmail, num, "300")

	return bool

}

func userDoc(name string, userEmail string, message string) string {
	html := `<div style="
    width: 100%;
    border-radius: 10px;
    box-shadow: 0 10px 10px rgba(0, 0, 0, 0.2);">
        <div style=" background-color: #2a265f; color: #fff;border-radius: 10px;display: flex;height: 40px;">
            <div style="letter-spacing: 1px;
            text-transform: uppercase;padding: 0 15px;line-height: 40px;">` + name + `:</div>
            <div style="letter-spacing: 1px;line-height: 40px;">` + userEmail + `</div>
        </div>
        <div style="text-indent: 2rem;letter-spacing: 1px;padding: 10px 15px;">
            ` + message + `
        </div>
    </div>`

	return html
}

func UserToMyEmail(name string, userEmail string, message string) bool {
	html := userDoc(name, userEmail, message)

	bool := initEmail("757909414@qq.com", html, "用户建议 Zjs-7579")

	//db.SetRedis(toEmail, num, "300")

	return bool
}
