package label

import (
	. "money-record/app/database"
	. "money-record/app/model/label"
	"money-record/app/util"

	"github.com/gin-gonic/gin"
)

func LabelRouter(router *gin.RouterGroup) {
	router.GET("/all", getAllLabel)
	router.POST("/add", addLabel)
	router.POST("/delete", deleteLabel)
	router.POST("/update", updateLabel)
}

// -------------- 标签列表 --------------

func getAllLabel(c *gin.Context) {
	utilGin := util.Gin{
		Ctx: c,
	}
	var msg LabelList
	rs := getRows(c)
	msg.List = rs
	utilGin.Response("Y", "success", msg)
}

func getRows(c *gin.Context) (label []Label) {
	DB.Table("label").Where("book_id = ? and type = ?", c.Query("bookId"), c.Query("type")).Find(&label)
	return
}

// -------------- 新增标签 --------------

func addLabel(c *gin.Context) {
	utilGin := util.Gin{
		Ctx: c,
	}
	var b Label
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
	utilGin.Response("N", "error", err)
}

// -------------- 删除标签 --------------

func deleteLabel(c *gin.Context) {
	utilGin := util.Gin{
		Ctx: c,
	}
	var b Label
	err := c.Bind(&b)
	if err != nil {
		utilGin.Response("N", "error", err)
		return
	}
	DB.Delete(&b)
	utilGin.Response("Y", "success", nil)
}

// -------------- 更新标签 --------------

func updateLabel(c *gin.Context) {
	utilGin := util.Gin{
		Ctx: c,
	}
	var b Label
	err := c.Bind(&b)
	if err != nil {
		utilGin.Response("N", "error", err)
		return
	}
	DB.Model(&b).Update("name", b.Name)
	utilGin.Response("Y", "success", nil)
}
