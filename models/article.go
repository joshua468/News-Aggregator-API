package models

type Article struct {
	ID     int    `gorm:"primaryKey" json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	URL    string `json:"url"`
}
