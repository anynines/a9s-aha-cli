package aha

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func HttpGet(url string, username string, password string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if len(username) > 0 {
		req.SetBasicAuth(username, password)
	}

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET %s returned status %d", url, resp.StatusCode)
	}

	return body, nil
}
