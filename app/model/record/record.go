package record

type Record struct {
	Id      int     `gorm:"primary_key" json:"id"`
	LabelId int     `json:"labelId"`
	Type    int     `json:"type"`
	Time    string  `json:"time"`
	Money   float32 `json:"money"`
	Remark  string  `json:"remark"`
}

type RecordQueryParams struct {
	current   int
	pageSize  int
	Type      int
	time      string
	startTime string
	endTime   string
	labelId   int
	money     string
	bookId    int
}

// TODO: 结构体继承和重写
type RecordListItemRes struct {
	Id     int     `gorm:"primary_key" json:"id"`
	Name   string  `json:"label"`
	Type   int     `json:"type"`
	Time   string  `json:"time"`
	Money  float32 `json:"money"`
	Remark string  `json:"remark"`
}

type TotalRes struct {
	Total       int     `json:"total"`
	TotalAmount float32 `json:"totalAmount"`
}

type RecordListRes struct {
	TotalRes
	List []RecordListItemRes `json:"list"`
}

// -------- bar 接口 ------------
type RecordBarList struct {
	Year  int     `json:"year"`
	Total float32 `json:"total"`
}
type RecordBarData struct {
	Income []RecordBarList `json:"income"`
	Pay    []RecordBarList `json:"pay"`
}

type RecordBarRes struct {
	List RecordBarData `json:"list"`
}

// -------- pie 接口 ------------
type RecordPieList struct {
	Total float32 `json:"totalAmount"`
	Name  string  `json:"labelName"`
}

type RecordPieRes struct {
	List []RecordPieList `json:"list"`
}
