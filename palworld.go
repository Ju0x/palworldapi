package palworldapi

import (
	"fmt"
	"strings"
	"time"
)

func New(host, username, password string) *PalworldAPI {

	if host == "" {
		host = "localhost:8212"
	}

	host = strings.TrimPrefix(host, "http://")
	host = strings.TrimPrefix(host, "https://")
	host = strings.TrimSuffix(host, "/")

	if username == "" {
		username = "admin"
	}

	return &PalworldAPI{
		Host:          host,
		Username:      username,
		AdminPassword: password,
	}
}

// Gets the server info
// Palworld docs: https://tech.palworldgame.com/api/rest-api/info
func (p *PalworldAPI) Info() (info *ServerInfo, err error) {
	err = p.get("/info", &info)

	if err != nil {
		return nil, fmt.Errorf("error getting server info: %w", err)
	}

	return info, err
}

// Gets the player list
// Palworld docs: https://tech.palworldgame.com/api/rest-api/players
func (p *PalworldAPI) Players() (players []*Player, err error) {
	var playerList *PlayerList

	err = p.get("/players", &playerList)

	if err != nil {
		return nil, fmt.Errorf("error getting player list: %w", err)
	}

	// This creates a list without content to prevent nil-pointer exceptions
	if playerList == nil {
		playerList = &PlayerList{
			Players: []*Player{},
		}
	}

	players = playerList.Players

	return
}

// Gets the server settings
// Palworld docs: https://tech.palworldgame.com/api/rest-api/settings
func (p *PalworldAPI) Settings() (settings *ServerSettings, err error) {
	err = p.get("/settings", &settings)

	if err != nil {
		return nil, fmt.Errorf("error getting server settings: %w", err)
	}
	return
}

// Gets the server metrics
// Palworld docs: https://tech.palworldgame.com/api/rest-api/metrics
func (p *PalworldAPI) Metrics() (metrics *ServerMetrics, err error) {
	err = p.get("/metrics", &metrics)

	if err != nil {
		return nil, fmt.Errorf("error getting server metrics: %w", err)
	}
	return
}

// Announces a message to the server
// Palworld docs: https://tech.palworldgame.com/api/rest-api/announce
func (p *PalworldAPI) Announce(message string) (err error) {
	err = p.post("/announce", Announce{Message: message})
	return
}

// Kicks a player from the server
// Palworld docs: https://tech.palworldgame.com/api/rest-api/kick
func (p *PalworldAPI) Kick(user_id string, message string) (err error) {
	err = p.post("/kick", KickPlayer{UserID: user_id, Message: message})
	return
}

// Bans a player from the server
// Palworld docs: https://tech.palworldgame.com/api/rest-api/ban
func (p *PalworldAPI) Ban(user_id string, message string) (err error) {
	err = p.post("/ban", BanPlayer{UserID: user_id, Message: message})
	return
}

// Unbans a player from the server
// Palworld docs: https://tech.palworldgame.com/api/rest-api/unban
func (p *PalworldAPI) Unban(user_id string) (err error) {
	err = p.post("/unban", UnbanPlayer{UserID: user_id})
	return
}

// Saves the world
// Palworld docs: https://tech.palworldgame.com/api/rest-api/save
func (p *PalworldAPI) SaveWorld() (err error) {
	err = p.post("/save", nil)
	return
}

// Shuts the server down
// Palworld docs: https://tech.palworldgame.com/api/rest-api/shutdown
func (p *PalworldAPI) Shutdown(waittime time.Duration, message string) (err error) {
	err = p.post("/shutdown", ServerShutdown{WaitTime: int(waittime.Seconds()), Message: message})
	return
}

// Force stops the server
// Palworld docs: https://tech.palworldgame.com/api/rest-api/stop
func (p *PalworldAPI) ForceStop() (err error) {
	err = p.post("/stop", nil)
	return
}
