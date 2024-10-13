package main

import (
	"fmt"
	"os"

	"github.com/Ju0x/palworldapi"
)

func main() {
	pal := palworldapi.New("http://localhost:8212", os.Getenv("USERNAME"), os.Getenv("ADMIN_PASSWORD"))

	// Get Server Info
	info, err := pal.Info()

	if err != nil {
		panic(err)
	}

	fmt.Println("Server Name:", info.ServerName)
	fmt.Println("Description:", info.Description)
	fmt.Println("Version:", info.Version)

	// Get Metrics
	metrics, err := pal.Metrics()

	if err != nil {
		panic(err)
	}

	fmt.Println()
	fmt.Println("--- Metrics ---")

	fmt.Println("FPS:", metrics.FPS)
	fmt.Println("Players Online:", metrics.CurrentPlayerNum)
	fmt.Println("Max. Players:", metrics.MaxPlayerNum)
	fmt.Println("Uptime (seconds):", metrics.Uptime)
}
