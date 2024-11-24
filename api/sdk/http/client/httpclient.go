package client

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

func DoRequest(url, method string, data interface{}) ([]byte, error) {
	client := http.Client{}
	ctx, cancel := context.WithTimeout(context.TODO(), 3*time.Second)
	defer cancel()
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(dataBytes))
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := json.Marshal(res.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}
