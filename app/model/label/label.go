package label

type Label struct {
	Id      int     `gorm:"primary_key" json:"id"`
	BookId      int     `json:"bookId"`
	Type      int     `json:"type"`
	Name    string  `json:"name"`
}