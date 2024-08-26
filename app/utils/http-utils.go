package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Get(url string, result interface{}) error {

	// Make the GET request
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return err
	}

	// Parse the JSON response
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return err
	}

	return nil
}
