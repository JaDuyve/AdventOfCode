package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func ReadHTTP(year, day int, session string) (string, error) {
	path := fmt.Sprintf("resources/day%d.txt", day)
	data, err := ReadFile(path)
	if err == nil {
		log.Println("Input already saved")
		return data, nil
	}

	if session == "" {
		return "", fmt.Errorf("session is required")
	}
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	session = fmt.Sprintf("session=%s", session)

	log.Printf("fetching input : %s", url)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("cookie", session)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(string(bodyBytes))
	}

	body := strings.Trim(string(bodyBytes), "\n")
	err = WriteFile(path, body)
	if err != nil {
		return "", err
	}

	return body, nil
}
