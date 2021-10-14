package record

type Record struct {
	Id      int     `gorm:"primary_key" json:"id"`
	LabelId		int 		`json:"labelId"`
	Type		int 		`json:"type"`
	Time    string  `json:"time"`
	Money		float32  `json:"money"`
	Remark    string  `json:"remark"`
}

type RecordQueryParams struct {
	current int
  pageSize int
  Type int
  time string
  startTime string
  endTime string
  labelId int
  money string
  bookId int
}

// TODO: 结构体继承和重写
type RecordListItemRes struct {
	Id      int     `gorm:"primary_key" json:"id"`
	Name		string 		`json:"labelName"`
	Type		int 		`json:"type"`
	Time    string  `json:"time"`
	Money		float32  `json:"money"`
	Remark    string  `json:"remark"`
}

type RecordListRes struct {
	Data []RecordListItemRes `json:"data"`
}

type RecordBarList struct {
	Year int `json:"year"`
	Total float32 `json:"total"`
}

type RecordBarData struct {
	Income []RecordBarList `json:"income"`
	Pay []RecordBarList `json:"pay"`
}

type RecordPieList struct {
	Total float32 `json:"total"`
	Name		string 		`json:"labelName"`
}