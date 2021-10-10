package book

import (
	"github.com/gin-gonic/gin"
	. "money-record/app/database"
	Book "money-record/app/model/book"
	"fmt"
)

func BookRouter(router *gin.RouterGroup) {
	router.GET("/all", getAllBook)
	router.POST("/add", addBook)
	router.POST("/delete", deleteBook)
	router.POST("/update", updateBook)
}



//应答体
type GormResponse struct {
	Code    string         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

var gormResponse GormResponse


func getAllBook(c *gin.Context) {
	rs := getRows()
	c.JSON(200, gin.H{
		"code": "Y",
		"result": rs,
	});
}

func getRows() (book []Book.BookModel)  {
	DB.Table("book").Find(&book)
	return;
}

// 新增账簿
func addBook(c *gin.Context) {
	var b Book.BookModel;
	err := c.Bind(&b)
	if err != nil {
		gormResponse.Message = "参数错误"
		gormResponse.Data = err
		c.JSON(200, gormResponse)
		return
	}
	
	tx := DB.Create(&b)
	fmt.Print(221, b, tx);
	if tx.RowsAffected > 0 {
		gormResponse.Code = "Y"
		gormResponse.Message = "写入成功"
		gormResponse.Data = "OK"
		c.JSON(200, gormResponse)
		return
	}
	//返回页面
	c.JSON(200, gin.H{
		"code": "N",
	})
}

// 删除账簿
func deleteBook(c *gin.Context) {
	var b Book.BookModel;
	err := c.Bind(&b)
	if err != nil {
		gormResponse.Message = "参数错误"
		gormResponse.Data = err
		c.JSON(200, gormResponse)
		return
	}
	DB.Delete(&b)
	c.JSON(200, gin.H{
		"code": "Y",
	});
}

// 更新账簿
func updateBook(c *gin.Context) {
	var b Book.BookModel;
	err := c.Bind(&b)
	if err != nil {
		gormResponse.Message = "参数错误"
		gormResponse.Data = err
		c.JSON(200, gormResponse)
		return
	}
	DB.Model(&b).Update("name", b.Name)
	c.JSON(200, gin.H{
		"code": "Y",
	});
}
