package record

type Record struct {
	Id      int     `gorm:"primary_key" json:"id"`
	Label		int 		`json:"label"`
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
  label int
  money string
  bookId int
}

type RecordBarList struct {
	Year int `json:"year"`
	Total float32 `json:"total"`
}

type RecordBarData struct {
	Income []RecordBarList `json:"income"`
	Pay []RecordBarList `json:"pay"`
}