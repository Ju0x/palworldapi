package main

import (
	"os"

	"github.com/Ju0x/palworldapi"
)

func main() {
	pal := palworldapi.New("http://localhost:8212", os.Getenv("USERNAME"), os.Getenv("ADMIN_PASSWORD"))

	// Gets the BanList whose URL is specified in the server config
	bannedPlayers, err := pal.BanList()

	if err != nil {
		panic(err)
	}

	for _, bannedId := range bannedPlayers {
		if bannedId == "steam_1234567890" {
			// Unbans a player from the server
			pal.Unban(bannedId)
		}
	}

	// Bans a player from the server
	pal.Ban("steam_0987654321", "Banned for killing Cremis")
}
