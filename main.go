package main

import (
	// "log"

	"github.com/rayfiyo/simpleLanguageModel/routers"
)

func main() {
	r := routers.InitRouter()
	r.Run()
}
