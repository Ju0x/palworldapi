package palworldapi

import "encoding/base64"

func (p *PalworldAPI) basicAuth() string {
	auth := p.Username + ":" + p.AdminPassword
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}
