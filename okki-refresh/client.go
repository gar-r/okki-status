package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var addr = ":12650"

func sendRefreshRequest(module string) error {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(fmt.Sprintf("http://localhost%v/invalidate/%v", addr, module))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("status code: %d, message: %s", resp.StatusCode, body)
	}
	return nil
}
