package models

type Response struct {
	Content string `json:"content"`
}

type Message struct {
	Mode    string `json:"mode"`
	Content string `json:"content"`
}

type storage struct {
	Query []string
	Reply []string
}

var Storage = storage{
	Query: []string{"あなたの名前は何ですか？"},
	Reply: []string{"はいそうです"},
}
