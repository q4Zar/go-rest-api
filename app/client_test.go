package main

import (
    "encoding/base64"
    "fmt"
    "io/ioutil"
    "net/http"
    "testing"
)

// BasicAuth generates the base64 encoded string for the basic auth header.
func BasicAuth(username, password string) string {
    auth := username + ":" + password
    return base64.StdEncoding.EncodeToString([]byte(auth))
}

// TestGetRouteWithBasicAuth tests a GET route with basic authentication.
func TestGetRouteWithBasicAuth(t *testing.T) {
    // Define the username, password, and the URL
    username := "admin"
    password := "admin"
    url := "http://127.0.0.1:8080/hello"

    // Create a new HTTP client
    client := &http.Client{}

    // Create a new request
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        t.Fatalf("Error creating request: %v", err)
    }

    // Set the Basic Authentication header
    auth := BasicAuth(username, password)
    req.Header.Add("Authorization", "Basic "+auth)

    // Perform the request
    resp, err := client.Do(req)
    if err != nil {
        t.Fatalf("Error performing request: %v", err)
    }
    defer resp.Body.Close()

    // Read the response body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        t.Fatalf("Error reading response body: %v", err)
    }

    // Print the response status and body
    fmt.Printf("Response Status: %s\n", resp.Status)
    fmt.Printf("Response Body: %s\n", string(body))

    // Add your assertions here
    // Example:
    if resp.StatusCode != http.StatusOK {
        t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
    }
}
