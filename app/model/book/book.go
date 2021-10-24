package book

type Book struct {
	Id   int    `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}

type BookList struct {
	List []Book `json:"list"`
}
