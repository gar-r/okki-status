package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

func sendRefreshRequest(module string) error {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(fmt.Sprintf("http://localhost%v/invalidate/%v", addr, module))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("invalid status code: " + string(resp.StatusCode))
	}
	return nil
}
