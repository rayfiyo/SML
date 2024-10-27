package models

type Response struct {
	Content string `json:"content"`
}

type Message struct {
	Mode    string `json:"mode"`
	Content string `json:"content"`
}

type Storage struct {
	Query []string
	Reply []string
}

var storage = Storage{
	Query: []string{"あなたの名前は何ですか？"},
	Reply: []string{"はいそうです"},
}
