package models

type TemplateData struct {
	CurrentYear     int
	Post            *Post
	Posts           []*Post
	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
	User            *User
}
