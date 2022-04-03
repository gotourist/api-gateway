package models

type Error struct {
	Field    string   `json:"field"`
	Messages []string `json:"messages"`
	Errors   []*Error `json:"errors"`
}
