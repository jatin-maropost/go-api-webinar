package models

// Book struct (Model) Api Model
type Book struct {
	ID string	`json:"id"`
	Isbn string	`json:"isbn"`
	Title string	`json:"title"`
	Author *Author	`json:"author"`
}