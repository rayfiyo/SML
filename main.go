package main

import (
	// "log"

	"github.com/rayfiyo/simple-language-model/routers"
)

func main() {
	r := routers.InitRouter()
	r.Run("localhost:8080")
}
