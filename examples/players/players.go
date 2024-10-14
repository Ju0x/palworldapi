package main

import (
	"fmt"
	"os"

	"github.com/Ju0x/palworldapi"
)

func main() {
	pal := palworldapi.New("http://localhost:8212", os.Getenv("USERNAME"), os.Getenv("ADMIN_PASSWORD"))

	players, err := pal.Players()

	if err != nil {
		panic(err)
	}

	// Can also be retrieved by ServerMetrics
	fmt.Printf("Players online: %d\n", len(players))

	if len(players) == 0 {
		return
	}

	for i, p := range players {
		fmt.Printf("[Nr. %d] Username: %s Steam Name: %s IP: %s Ping: %fms Level: %d Location: (X: %f, Y: %f)\n", i+1, p.Name, p.AccountName, p.IP, p.Ping, p.Level, p.LocationX, p.LocationY)
	}
}
