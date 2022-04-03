package models

type UpdatePostRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
