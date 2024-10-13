package main

import (
	"os"
	"time"

	"github.com/Ju0x/palworldapi"
)

func main() {
	pal := palworldapi.New("http://localhost:8212", os.Getenv("USERNAME"), os.Getenv("ADMIN_PASSWORD"))

	pal.Announce("Hello World!")

	pal.SaveWorld()

	pal.Announce("Server will shutdown soon!")

	players, err := pal.Players()

	if err != nil {
		panic(err)
	}

	pal.Shutdown(30*6, "Server shuts down in 60 Seconds!")

	time.Sleep(5 * time.Second)
	if len(players) > 0 {
		pal.Kick(players[0].UserID, "Kicked just for fun!")
	}
}
