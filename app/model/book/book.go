package book

type BookModel struct {
	Id      int     `gorm:"primary_key" json:"id"`
	Name    string  `json:"name"`
}