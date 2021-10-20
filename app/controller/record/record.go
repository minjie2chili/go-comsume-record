package record

import (
	"github.com/gin-gonic/gin"
	. "money-record/app/database"
	. "money-record/app/model/record"
	"gorm.io/gorm"
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

func getResultByCondition(c *gin.Context, tx1 *gorm.DB)(db *gorm.DB) {
	tx := tx1
	// GetQuery返回两个参数，第一个是参数值，第二个参数是参数是否存在的bool值，可以用来判断参数是否存在
	Type, typeExist := c.GetQuery("type")
	label, labelExist := c.GetQuery("label")
	startTime, startTimeExist := c.GetQuery("startTime")
	endTime, endTimeExist := c.GetQuery("endTime")
	money := c.Query("money")
	time := c.Query("time")

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
	return tx
}

func getAllRecord(c *gin.Context) {
	rs, total := getRows(c)
	var msg RecordListRes;
	msg.Data = rs
	msg.Total = total.Total
	msg.TotalAmount = total.TotalAmount

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
	Type := c.Query("type")
	startTime, startTimeExist := c.GetQuery("startTime")
	endTime, endTimeExist := c.GetQuery("endTime")
	var s1 string;
	if startTimeExist && endTimeExist {
		s1 = "and record.time between '" + startTime + "' and '" + endTime + "'";
	}
	tx := DB.Table("record").Raw("select name, sum(money) as total from (select * from label where type = ?) a join record on record.label_id = a.id and record.type = ? "+ s1 +" group by record.label_id", Type, Type)
	tx.Scan(&record)
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

func getRows(c *gin.Context)(res []RecordListItemRes, total TotalRes)  {
	var b RecordQueryParams;
	c.Bind(&b)
	// 条件查询
	bookId := c.Query("bookId")
	// 分页参数
	page,_ := strconv.Atoi(c.DefaultQuery("page", "0"))
	pageSize,_ := strconv.Atoi(c.Query("pageSize"))
	offset := (page - 1) * pageSize
	
	tx := DB.Table("record").Select("record.*, label.name").Joins("join label on record.label_id = label.id and label.book_id = ?", bookId).Offset(offset).Limit(pageSize)
	tx2 := DB.Table("record").Select("count(*) as total, SUM(money) as TotalAmount").Joins("join label on record.label_id = label.id and label.book_id = ?", bookId)
	fmt.Println(333, tx2)
	tx = getResultByCondition(c, tx)
	tx2 = getResultByCondition(c, tx2)
	tx.Find(&res)
	tx2.Scan(&total)
	fmt.Println(222, total)
	if tx != nil {
		fmt.Println(222, tx2)
	}
	return res, total;
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
