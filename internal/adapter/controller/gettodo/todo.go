package gettodo

type Todo struct {
	Id       int    `json:"id"`
	UserId   int    `json:"userId"`
	Content  string `json:"content"`
	MediaUrl string `json:"mediaUrl,omitempty"`
}
