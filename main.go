package main

import (
	"fmt"
	"net/http"
	"leaderboard-server/leaderboard"
)

const port = 8080

func main() {
	err := leaderboard.Init()
	if err != nil {
		panic(err)
	}

	leaderboard.Bind()

	fmt.Printf("Listening on port %d...\n", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		panic(err)
	}
}
