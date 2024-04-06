package request

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func MakeRequest(method, url, token string) (interface{}, error) {
	result := make(map[string]interface{})
	client := http.DefaultClient
	req, _ := http.NewRequest(method, url, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Api-Token", token)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func MakeRequestWithBody(method, url, token string, p interface{}) (string, error) {
	client := http.DefaultClient
	data, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(data))

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Api-Token", token)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	return string(body), nil

}
