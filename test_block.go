package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env") // loads .env file
	githubToken := os.Getenv("GITHUB_TOKEN")
	repo := "tesing-purpose/testing-blocking-unblocking"


	url := fmt.Sprintf("https://api.github.com/repos/%s/statuses/%s", repo, commitSHA)

	payload := map[string]string{
		"state":       "failure",
		"context":     "pipeline-block",
		"description": "Repo blocked: testing parity drift blocking",
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+githubToken)
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Response: %s\n", string(respBody))
}
