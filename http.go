package palworldapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	Endpoint = "/v1/api"
)

func (p *PalworldAPI) headers(r *http.Request) {
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", p.basicAuth())
}

func (p *PalworldAPI) apiUrl(s string) string {
	return "http://" + p.Host + Endpoint + s
}

func (p *PalworldAPI) get(path string, result interface{}) error {
	url := p.apiUrl(path)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	p.headers(req)

	res, err := p.client.Do(req)
	if err != nil {
		return fmt.Errorf("error executing request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP status code is %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	err = json.Unmarshal(body, result)
	if err != nil {
		return fmt.Errorf("error unmarshalling response: %w", err)
	}

	return nil
}

func (p *PalworldAPI) post(path string, body interface{}) error {
	url := p.apiUrl(path)

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("error marshalling body: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	p.headers(req)

	res, err := p.client.Do(req)
	if err != nil {
		return fmt.Errorf("error executing request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP status code is %d", res.StatusCode)
	}

	return nil
}
