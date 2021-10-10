package route

import (
	"money-record/app/controller/book"
	"money-record/app/controller/label"
	"money-record/app/controller/record"
	"github.com/gin-gonic/gin"
)

func CollectRoute(router *gin.Engine) {

	bookGroup := router.Group("/book")
	book.BookRouter(bookGroup)
	// 标签的路由
	labelGroup := router.Group("/label")
	label.LabelRouter(labelGroup)
 
	recordGroup := router.Group("/record")
	record.RecordRouter(recordGroup)
}
