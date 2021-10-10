package label

import (
	"github.com/gin-gonic/gin"
	. "money-record/app/database"
	"fmt"
)

func LabelRouter(router *gin.RouterGroup) {
	router.GET("/all", getAllLabel)
	router.POST("/add", addLabel)
	router.POST("/delete", deleteLabel)
	router.POST("/update", updateLabel)
}

type Label struct {
	Id      int     `gorm:"primary_key" json:"id"`
	Name    string  `json:"name"`
}

//应答体
type GormResponse struct {
	Code    string         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

var gormResponse GormResponse


func getAllLabel(c *gin.Context) {
	rs := getRows()
	c.JSON(200, gin.H{
		"code": "Y",
		"result": rs,
	});
}

func getRows() (label []Label)  {
	DB.Table("label").Find(&label)
	return;
}

// 新增标签
func addLabel(c *gin.Context) {
	var b Label;
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

// 删除标签
func deleteLabel(c *gin.Context) {
	var b Label;
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

// 更新标签
func updateLabel(c *gin.Context) {
	var b Label;
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
