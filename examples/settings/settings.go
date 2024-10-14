package main

import (
	"fmt"
	"os"

	"github.com/Ju0x/palworldapi"
)

func main() {
	pal := palworldapi.New("http://localhost:8212", os.Getenv("USERNAME"), os.Getenv("ADMIN_PASSWORD"))

	settings, err := pal.Settings()

	if err != nil {
		panic(err)
	}

	// Some example values you can retrieve
	fmt.Println("Server Name: ", settings.ServerName)
	fmt.Println("Rest API Port: ", settings.RESTAPIPort)
	fmt.Println("PVP Enabled: ", settings.IsPvP)
	fmt.Println("Banlist URL: ", settings.BanListURL)
}
