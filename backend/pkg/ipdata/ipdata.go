package ipdata

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	apiURL  = "https://ipwho.is/"
	timeout = 2 * time.Second
)

type IPData struct {
	City    string `json:"city"`
	Country string `json:"country"`
	Region  string `json:"region"`
}

type apiResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"message,omitempty"`
	IPData
}

func GetIPWho(ctx context.Context, ip string) (*IPData, error) {
	client := &http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequestWithContext(ctx, "GET", apiURL+ip, nil)
	if err != nil {
		return nil, fmt.Errorf("request creation failed: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("API request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		switch resp.StatusCode {
		case http.StatusTooManyRequests:
			return nil, errors.New("API rate limit exceeded")
		case http.StatusBadRequest:
			return nil, errors.New("invalid IP address format")
		default:
			return nil, fmt.Errorf("unexpected API status: %d", resp.StatusCode)
		}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("response read failed: %w", err)
	}

	var result apiResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("JSON unmarshal failed: %w", err)
	}

	if !result.Success {
		if result.Error != "" {
			return nil, fmt.Errorf("ipwhois error: %s", result.Error)
		}
		return nil, errors.New("failed to get IP data")
	}

	return &result.IPData, nil
}
