package book

import (
	. "money-record/app/database"
	. "money-record/app/model/book"
	"money-record/app/util"

	"github.com/gin-gonic/gin"
)

func BookRouter(router *gin.RouterGroup) {
	router.GET("/all", getAllBook)
	router.POST("/add", addBook)
	router.POST("/delete", deleteBook)
	router.POST("/update", updateBook)
}

// -------------- 账簿列表 --------------

func getAllBook(c *gin.Context) {
	utilGin := util.Gin{
		Ctx: c,
	}
	var res BookList
	rs := getRows()
	res.List = rs
	utilGin.Response("Y", "success", res)
}

func getRows() (book []Book) {
	DB.Table("book").Find(&book)
	return
}

// -------------- 新增账簿 --------------

func addBook(c *gin.Context) {
	utilGin := util.Gin{
		Ctx: c,
	}
	var b Book
	err := c.Bind(&b)
	if err != nil {
		utilGin.Response("N", "error", err)
		return
	}

	res := DB.Create(&b)

	if res.RowsAffected > 0 {
		utilGin.Response("Y", "success", nil)
		return
	}
	utilGin.Response("N", "error", nil)
}

// -------------- 删除账簿 --------------

func deleteBook(c *gin.Context) {
	utilGin := util.Gin{
		Ctx: c,
	}
	var b Book
	err := c.Bind(&b)
	if err != nil {
		utilGin.Response("N", "error", err)
		return
	}
	DB.Delete(&b)
	utilGin.Response("Y", "success", nil)
}

// -------------- 更新账簿 --------------

func updateBook(c *gin.Context) {
	utilGin := util.Gin{
		Ctx: c,
	}
	var b Book
	err := c.Bind(&b)
	if err != nil {
		utilGin.Response("N", "error", err)
		return
	}
	DB.Model(&b).Update("name", b.Name)
	utilGin.Response("Y", "success", nil)
}
