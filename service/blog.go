package service

import (
	"crypto/md5"
	"fmt"
	"indexServer/db"
	"indexServer/logger"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Tag_Class struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Cerate string `json:"cerate"`
}

//get
func get(sql string) ([]Tag_Class, error) {
	var list = make([]Tag_Class, 0, 100)

	//sql := "select id,name,cerate from Tag;"
	rows, err := db.BlogMysqlDB.Query(sql)
	//fmt.Println(rows)
	if err != nil {
		logger.SetLogger(1, sql+": Mysql查询出错")
		log.Println(sql + ":Mysql查询出错")
		return nil, err
	}

	for rows.Next() {
		var data Tag_Class
		err = rows.Scan(&data.Id, &data.Name, &data.Cerate)
		if err != nil {
			logger.SetLogger(1, sql+"绑定数据失败")
			log.Println(sql + "绑定数据失败")
			return nil, err
		}
		list = append(list, data)
	}

	return list, nil

}

//add
func add(ctx *gin.Context, sql string) (int64, error) {
	var name Tag_Class
	if err := ctx.ShouldBindJSON(&name); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		logger.SetLogger(1, "解析json数据出错")
		ctx.JSON(http.StatusOK, gin.H{
			"code":  200,
			"error": err.Error(),
		})
		return -1, err
	}
	//fmt.Printf(name.Name)

	//sql := "insert into Tag(name) values(?);"
	res, err := db.BlogMysqlDB.Exec(sql, name.Name)
	if err != nil {
		logger.SetLogger(1, sql+"绑定数据失败")
		log.Println(sql + "绑定数据失败")
		return -1, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		logger.SetLogger(1, "获取长度失败")
		log.Println("获取长度失败")
		return -1, err
	}

	return id, nil
}

//del
func del(ctx *gin.Context, sql string) error {
	var id Tag_Class
	if err := ctx.ShouldBindJSON(&id); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		logger.SetLogger(1, "解析json数据出错")
		ctx.JSON(http.StatusOK, gin.H{
			"code":  200,
			"error": err.Error(),
		})
		return err
	}
	//fmt.Printf(name.Name)

	_, err := db.BlogMysqlDB.Exec(sql, id.Id)
	if err != nil {
		logger.SetLogger(1, sql+"绑定数据失败")
		log.Println(sql + "绑定数据失败")
		return err
	}
	return nil
}

//标签------------------------------------------------

func GetTag(ctx *gin.Context) {

	//var list = make([]Tag, 0, 100)

	sql := "select id,name,cerate from Tag;"

	list, err := get(sql)
	if err != nil {
		log.Println(sql + "获取数据失败")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": list,
	})
}

func AddTag(ctx *gin.Context) {

	sql := "insert into Tag(name) values(?);"

	id, err := add(ctx, sql)
	if err != nil {
		log.Println(sql + "添加数据失败")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"id":   id,
		"msg":  "标签添加成功",
	})
}

func DelTag(ctx *gin.Context) {

	sql := "delete from Tag where id = ?;"

	err := del(ctx, sql)
	if err != nil {
		log.Println(sql + "删除数据失败")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "标签删除成功",
	})
}

//分类------------------------------------------------

func GetClass(ctx *gin.Context) {

	sql := "select id,name,cerate from Class;"
	list, err := get(sql)
	if err != nil {
		log.Println(sql + "获取数据失败")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": list,
	})
}

func AddClass(ctx *gin.Context) {

	sql := "insert into Class(name) values(?);"
	id, err := add(ctx, sql)
	if err != nil {
		log.Println(sql + "添加数据失败")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"id":   id,
		"msg":  "分类添加成功",
	})
}

func DelClass(ctx *gin.Context) {

	sql := "delete from Class where id = ?;"
	err := del(ctx, sql)
	if err != nil {
		log.Println(sql + "删除数据失败")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "分类删除成功",
	})
}

//Markdown--------------------------------------------

type Markdown struct {
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
	Id     int    `json:"id"`
	Uid    string `json:"uid"`
	Title  string `json:"title"`
	Class  string `json:"class"`
	Tag    string `json:"tag"`
	Text   string `json:"text"`
	Cerate string `json:"cerate"`
}

func GetMarkdown(ctx *gin.Context) {

	var list []Markdown

	sql := "select id,uid,title,class,tag,text,cerate from Blog;"
	rows, err := db.BlogMysqlDB.Query(sql)
	//fmt.Println(rows)
	if err != nil {
		logger.SetLogger(1, sql+": Mysql查询出错")
		log.Println(sql + ":Mysql查询出错")
		return
	}

	for rows.Next() {
		var data Markdown
		err = rows.Scan(&data.Id, &data.Uid, &data.Title, &data.Class, &data.Tag, &data.Text, &data.Cerate)
		if err != nil {
			logger.SetLogger(1, sql+"绑定数据失败")
			log.Println(sql + "绑定数据失败")
			return
		}

		list = append(list, data)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": list,
	})
}

func AddMarkdown(ctx *gin.Context) {
	var markdown Markdown
	if err := ctx.ShouldBindJSON(&markdown); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		logger.SetLogger(1, "解析json数据出错")
		ctx.JSON(http.StatusOK, gin.H{
			"code":  200,
			"error": err.Error(),
		})
	}
	//fmt.Printf(name.Name)
	var data = []byte(markdown.Title)
	uid := fmt.Sprintf("%x", md5.Sum(data))
	fmt.Println(uid, markdown)
	sql := "insert into Blog(uid,title,class,tag,text) values(?,?,?,?,?);"
	_, err := db.BlogMysqlDB.Exec(sql, uid, markdown.Title, markdown.Class, markdown.Tag, markdown.Text)
	if err != nil {
		logger.SetLogger(1, sql+"数据添加失败")
		log.Println(sql + "数据添加失败")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "文章添加成功",
	})
}

func PutMarkdown(ctx *gin.Context) {

}

func DelMarkdown(ctx *gin.Context) {
	var markdown Markdown
	if err := ctx.ShouldBindJSON(&markdown); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		logger.SetLogger(1, "解析json数据出错")
		ctx.JSON(http.StatusOK, gin.H{
			"code":  200,
			"error": err.Error(),
		})
	}
	//fmt.Printf(name.Name)
	sql := "delete from Blog where uid = ?;"
	_, err := db.BlogMysqlDB.Exec(sql, markdown.Uid)
	if err != nil {
		logger.SetLogger(1, sql+"数据删除失败")
		log.Println(sql + "数据删除失败")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "文章删除成功",
	})
}
