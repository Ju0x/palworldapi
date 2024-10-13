package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Ju0x/palworldapi"
)

func main() {
	pal := palworldapi.New("http://localhost:8212", os.Getenv("USERNAME"), os.Getenv("ADMIN_PASSWORD"))

	players, err := pal.Players()

	if err != nil {
		panic(err)
	}

	// Can also be retreived by ServerMetrics
	fmt.Printf("Players online: %d\n", len(players))

	if len(players) == 0 {
		return
	}

	bannedPlayers := make([]*palworldapi.Player, 0)

	for i, p := range players {
		if p.Name == "Ju0x" {
			// Ban if Ju0x is on the Server
			err := pal.Ban(p.UserID, "Has Ju0x as ingame Name")

			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Printf("Banned %s (%s) with User-ID %s\n", p.Name, p.AccountName, p.UserID)
			bannedPlayers = append(bannedPlayers, p)
			continue
		}

		fmt.Printf("[Nr. %d] Username: %s Steam Name: %s IP: %s Ping: %fms Level: %d Location: (X: %f, Y: %f)\n", i+1, p.Name, p.AccountName, p.IP, p.Ping, p.Level, p.LocationX, p.LocationY)
	}

	// Unban users with the name Ju0x after 2 Minutes
	time.Sleep(2 * time.Minute)

	for _, p := range bannedPlayers {
		err := pal.Unban(p.UserID)

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		fmt.Printf("Unbanned %s (%s) with User-ID %s\n", p.Name, p.AccountName, p.UserID)
	}
}
