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

func (p *PalworldAPI) request(method string, path string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, p.apiUrl(path), bytes.NewBuffer(body))

	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	p.headers(req)

	res, err := p.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing request: %w", err)
	}

	if res.StatusCode != http.StatusOK {

		if res.StatusCode == http.StatusUnauthorized {
			return nil, fmt.Errorf("401: please set correct username and password - for more information read the api docs: https://tech.palworldgame.com/")
		}

		return nil, fmt.Errorf("http status code is %d", res.StatusCode)
	}
	return res, err
}

// Executes a GET request to the REST API
// Response data will be written in the result param
func (p *PalworldAPI) get(path string, result interface{}) error {
	res, err := p.request("GET", path, nil)

	if err != nil {
		return fmt.Errorf("GET request failed: %w", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	res.Body.Close()

	err = json.Unmarshal(body, result)
	if err != nil {
		return fmt.Errorf("error unmarshalling response: %w", err)
	}

	return nil
}

// Executes a POST request to the REST API
// Palworlds Server REST API doesn't provide relevant result data for POST requests other than the status code
func (p *PalworldAPI) post(path string, body interface{}) error {
	jsonBody, err := json.Marshal(body)

	if err != nil {
		return fmt.Errorf("error marshalling body: %w", err)
	}

	res, err := p.request("POST", path, jsonBody)

	if err != nil {
		return fmt.Errorf("POST request failed: %w", err)
	}

	res.Body.Close()

	return nil
}
