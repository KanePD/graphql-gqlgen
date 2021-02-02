package model

type Article struct {
	ID     	string `json:"id"`
	Title   string `json:"Title"`
	Content string `json:"Content"`
	Author 	User   `json:"user"`
	Section Section `json:"secion"`
}
