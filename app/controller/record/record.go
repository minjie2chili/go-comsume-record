package record

import (
	"github.com/gin-gonic/gin"
	. "money-record/app/database"
	. "money-record/app/model/record"
	"strconv"
	"fmt"
)

func RecordRouter(router *gin.RouterGroup) {
	router.GET("/list", getAllRecord)
	router.GET("/pie/list", getPieRecord)
	router.GET("/bar/list", getBarRecord)
	router.POST("/add", addRecord)
	router.POST("/delete", deleteRecord)
	router.POST("/update", updateRecord)
}

//应答体
type GormResponse struct {
	Code    string         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

var gormResponse GormResponse


func getAllRecord(c *gin.Context) {
	rs := getRows(c)
	var msg RecordListRes;
	msg.Data = rs
	c.JSON(200, gin.H{
		"code": "Y",
		"result": msg,
	});
}

func getPieRecord(c *gin.Context) {
	rs := getPieRows(c)
	c.JSON(200, gin.H{
		"code": "Y",
		"result": rs,
	});
}

func getPieRows(c *gin.Context) (record []RecordPieList)  {
	// var b RecordQueryParams;
	// c.Bind(&b)
	// // 条件查询
	// Type, typeExist := c.GetQuery("type")
	// startTime, startTimeExist := c.GetQuery("startTime")
	// endTime, endTimeExist := c.GetQuery("endTime")
	DB.Table("record").Raw("select name, sum(money) as total from (select * from label where type = 1) a join record on record.label_id = a.id and record.type = 1 group by record.label_id;").Scan(&record)
	return
}

func getBarRecord(c *gin.Context) {
	rs := getBarRows(c)
	c.JSON(200, gin.H{
		"code": "Y",
		"result": rs,
	});
}

func getBarRows(c *gin.Context) (record RecordBarData)  {
	record.Pay = make([]RecordBarList, 0)
	// 通过Scan方法把sql返回的数据放入我们的结构体中
	DB.Table("record").Raw("SELECT DATE_FORMAT(time,'%Y') year ,SUM(money) total from record where type = 1 group by year").Scan(&record.Income)
	DB.Table("record").Raw("SELECT DATE_FORMAT(time,'%Y') year ,SUM(money) total from record where type = 2 group by year").Scan(&record.Pay)
	return
}

func getRows(c *gin.Context) (record []RecordListItemRes)  {
	var b RecordQueryParams;
	c.Bind(&b)
	// 条件查询
	bookId := c.Query("bookId")
	// GetQuery返回两个参数，第一个是参数值，第二个参数是参数是否存在的bool值，可以用来判断参数是否存在
	Type, typeExist := c.GetQuery("type")
	label, labelExist := c.GetQuery("label")
	startTime, startTimeExist := c.GetQuery("startTime")
	endTime, endTimeExist := c.GetQuery("endTime")
	money := c.Query("money")
	time := c.Query("time")
	// 分页参数
	page,_ := strconv.Atoi(c.DefaultQuery("page", "0"))
	pageSize,_ := strconv.Atoi(c.Query("pageSize"))
	offset := (page - 1) * pageSize
	
	tx := DB.Table("record").Select("record.*, label.name").Joins("join label on record.label_id = label.id and label.book_id = ?", bookId).Offset(offset).Limit(pageSize)

	if typeExist {
		tx = tx.Where("record.type = ?", Type)
	}
	if labelExist {
		tx = tx.Where("record.label_id = ?", label)
	}
	if startTimeExist && endTimeExist {
		tx = tx.Where("record.time between ? and ?", startTime, endTime)
	}
	// 处理金额升降序
	if money == "ascend" {
		tx = tx.Order("money")
	} else if money == "descend" {
		
		tx = tx.Order("money desc")
	}
	// 处理type升降序
	if time == "ascend" {
		tx = tx.Order("time")
	} else {
		tx = tx.Order("time desc")
	}
	tx.Find(&record)
	if tx != nil {
		fmt.Println(tx)
	}
	return;
}

// 新增记录
func addRecord(c *gin.Context) {
	var b Record;
	err := c.Bind(&b)
	
	if err != nil {
		gormResponse.Message = "参数错误"
		gormResponse.Data = err
		c.JSON(200, gormResponse)
		return
	}
	
	tx := DB.Create(&b)
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

// 删除记录
func deleteRecord(c *gin.Context) {
	var b Record;
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

// 更新记录
func updateRecord(c *gin.Context) {
	var b Record;
	err := c.Bind(&b)
	if err != nil {
		gormResponse.Message = "参数错误"
		gormResponse.Data = err
		c.JSON(200, gormResponse)
		return
	}
	tx := DB.Model(&b).Updates(&b)
	if tx.RowsAffected > 0 {
		gormResponse.Code = "Y"
		gormResponse.Message = "写入成功"
		gormResponse.Data = "OK"
		c.JSON(200, gormResponse)
		return
	}
	c.JSON(200, gin.H{
		"code": "Y",
		"msg": tx,
	});
}
