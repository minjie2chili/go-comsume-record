package record

import (
	. "money-record/app/database"
	. "money-record/app/model/record"
	"money-record/app/util"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RecordRouter(router *gin.RouterGroup) {
	router.GET("/list", getAllRecord)
	router.GET("/pie/list", getPieRecord)
	router.GET("/bar/list", getBarRecord)
	router.POST("/add", addRecord)
	router.POST("/delete", deleteRecord)
	router.POST("/update", updateRecord)
}

func getResultByCondition(c *gin.Context, originRes *gorm.DB) (db *gorm.DB) {
	res := originRes
	// GetQuery返回两个参数，第一个是参数值，第二个参数是参数是否存在的bool值，可以用来判断参数是否存在
	Type, typeExist := c.GetQuery("type")
	label, labelExist := c.GetQuery("label")
	startTime, startTimeExist := c.GetQuery("startTime")
	endTime, endTimeExist := c.GetQuery("endTime")
	money := c.Query("money")
	time := c.Query("time")

	if typeExist {
		res = res.Where("record.type = ?", Type)
	}
	if labelExist {
		res = res.Where("record.label_id = ?", label)
	}
	if startTimeExist && endTimeExist {
		res = res.Where("record.time between ? and ?", startTime, endTime)
	}
	// 处理金额升降序
	if money == "ascend" {
		res = res.Order("money")
	} else if money == "descend" {

		res = res.Order("money desc")
	}
	// 处理type升降序
	if time == "ascend" {
		res = res.Order("time")
	} else {
		res = res.Order("time desc")
	}
	return res
}

func getAllRecord(c *gin.Context) {
	utilGin := util.Gin{
		Ctx: c,
	}
	rs, total := getRows(c)
	var msg RecordListRes
	msg.List = rs
	msg.Total = total.Total
	msg.TotalAmount = total.TotalAmount

	utilGin.Response("Y", "success", msg)
}

// -------- pie 接口 ------------

func getPieRecord(c *gin.Context) {
	utilGin := util.Gin{
		Ctx: c,
	}
	rs := getPieRows(c)
	var res RecordPieRes
	res.List = rs
	utilGin.Response("Y", "success", res)
}

func getPieRows(c *gin.Context) (record []RecordPieList) {
	Type := c.Query("type")
	startTime, startTimeExist := c.GetQuery("startTime")
	endTime, endTimeExist := c.GetQuery("endTime")
	var s1 string
	if startTimeExist && endTimeExist {
		s1 = "and record.time between '" + startTime + "' and '" + endTime + "'"
	}
	res := DB.Table("record").Raw("select name, sum(money) as total from (select * from label where type = ?) a join record on record.label_id = a.id and record.type = ? "+s1+" group by record.label_id", Type, Type)
	res.Scan(&record)
	return
}

// -------- bar 接口 ------------

func getBarRecord(c *gin.Context) {
	utilGin := util.Gin{
		Ctx: c,
	}
	rs := getBarRows(c)
	utilGin.Response("Y", "success", rs)
}

func getBarRows(c *gin.Context) (record RecordBarData) {
	record.Pay = make([]RecordBarList, 0)
	// 通过Scan方法把sql返回的数据放入我们的结构体中
	DB.Table("record").Raw("SELECT DATE_FORMAT(time,'%Y') year ,SUM(money) total from record where type = 1 group by year").Scan(&record.Income)
	DB.Table("record").Raw("SELECT DATE_FORMAT(time,'%Y') year ,SUM(money) total from record where type = 2 group by year").Scan(&record.Pay)
	return
}

// 获取所以记录
// TODO: TotalRes是RecordListItemRes的一部分
func getRows(c *gin.Context) (res []RecordListItemRes, total TotalRes) {
	var b RecordQueryParams
	c.Bind(&b)
	// 条件查询
	bookId := c.Query("bookId")
	// 分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	offset := (page - 1) * pageSize

	listRes := DB.Table("record").Select("record.*, label.name").Joins("join label on record.label_id = label.id and label.book_id = ?", bookId).Offset(offset).Limit(pageSize)
	totalRes := DB.Table("record").Select("count(*) as total, SUM(money) as TotalAmount").Joins("join label on record.label_id = label.id and label.book_id = ?", bookId)

	listRes = getResultByCondition(c, listRes)
	totalRes = getResultByCondition(c, totalRes)

	listRes.Find(&res)
	totalRes.Scan(&total)

	return res, total
}

// 新增记录

func addRecord(c *gin.Context) {
	utilGin := util.Gin{
		Ctx: c,
	}
	var b Record
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
	utilGin.Response("N", "error", res)
}

// 删除记录
func deleteRecord(c *gin.Context) {
	utilGin := util.Gin{
		Ctx: c,
	}
	var b Record
	err := c.Bind(&b)
	if err != nil {
		utilGin.Response("N", "error", nil)
		return
	}
	DB.Delete(&b)
	utilGin.Response("Y", "success", nil)
}

// 更新记录
func updateRecord(c *gin.Context) {
	utilGin := util.Gin{
		Ctx: c,
	}
	var b Record
	err := c.Bind(&b)
	if err != nil {
		utilGin.Response("N", "error", err)
		return
	}
	res := DB.Model(&b).Updates(&b)
	if res.RowsAffected > 0 {
		utilGin.Response("Y", "success", nil)
		return
	}
	utilGin.Response("N", "error", res)
}
