package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var (
	apiURL string
	reader *bufio.Reader
	token  string
)

func main() {
	reader = bufio.NewReader(os.Stdin)
	fmt.Print("Enter the API URL: ")
	apiURL, _ = reader.ReadString('\n')
	apiURL = strings.TrimSpace(apiURL)

	for {
		fmt.Println("Which function do you want to call?")
		fmt.Println("1: Create user")
		fmt.Println("2: Log in user")
		fmt.Println("3: Create EUR asset")
		fmt.Println("4: Create USD asset")
		fmt.Println("5: Get user assets")
		fmt.Println("6: Create order")
		fmt.Println("7: Check for orders")
		fmt.Println("8: Check for new balances")
		fmt.Println("9: Exit")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			createUser()
		case "2":
			logInUser()
		case "3":
			createAsset("EUR")
		case "4":
			createAsset("USD")
		case "5":
			getUserAssets()
		case "6":
			createOrder()
		case "7":
			checkForOrders()
		case "8":
			checkForNewBalances()
		case "9":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func createUser() {
	fmt.Println("Creating user")
	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)
	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	userPayload := map[string]string{"username": username, "password": password}
	userPayloadBytes, _ := json.Marshal(userPayload)
	resp, err := http.Post(fmt.Sprintf("%s/users", apiURL), "application/json", bytes.NewBuffer(userPayloadBytes))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(resp)
	defer resp.Body.Close()
	fmt.Println("User created")
}

func logInUser() {
	fmt.Println("Logging in user")
	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)
	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	userPayload := map[string]string{"username": username, "password": password}
	userPayloadBytes, _ := json.Marshal(userPayload)
	resp, err := http.Post(fmt.Sprintf("%s/login", apiURL), "application/json", bytes.NewBuffer(userPayloadBytes))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var loginResponse map[string]interface{}
	json.Unmarshal(body, &loginResponse)
	if t, ok := loginResponse["token"].(string); ok {
		token = t
		fmt.Printf("Token: %s\n", token)
	} else {
		fmt.Println("Failed to log in")
	}
}

func createAsset(assetType string) {
	if token == "" {
		fmt.Println("You need to log in first.")
		return
	}
	fmt.Printf("Creating %s asset\n", assetType)
	fmt.Printf("Enter %s balance: ", assetType)
	var balance float64
	fmt.Scanf("%f\n", &balance)

	assetPayload := map[string]interface{}{"assetType": assetType, "balance": balance}
	assetPayloadBytes, _ := json.Marshal(assetPayload)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/assets", apiURL), bytes.NewBuffer(assetPayloadBytes))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()
	fmt.Printf("Asset %s created\n", assetType)
}

func getUserAssets() {
	if token == "" {
		fmt.Println("You need to log in first.")
		return
	}
	fmt.Println("Getting user assets")

	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/assets?fields=balance,asset_type,user_id", apiURL), nil)
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func createOrder() {
	if token == "" {
		fmt.Println("You need to log in first.")
		return
	}
	fmt.Println("Creating order")

	fmt.Print("Enter amount: ")
	var amount float64
	fmt.Scanf("%f\n", &amount)
	fmt.Print("Enter price: ")
	var price float64
	fmt.Scanf("%f\n", &price)
	fmt.Print("Enter side (BUY/SELL): ")
	side, _ := reader.ReadString('\n')
	side = strings.TrimSpace(side)
	fmt.Print("Enter asset pair (e.g., USD-EUR): ")
	assetPair, _ := reader.ReadString('\n')
	assetPair = strings.TrimSpace(assetPair)

	orderPayload := map[string]interface{}{"amount": amount, "price": price, "side": side, "assetPair": assetPair}
	orderPayloadBytes, _ := json.Marshal(orderPayload)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/orders", apiURL), bytes.NewBuffer(orderPayloadBytes))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Order created")
}

func checkForOrders() {
	if token == "" {
		fmt.Println("You need to log in first.")
		return
	}
	fmt.Println("Checking for orders")

	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/orders", apiURL), nil)
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func checkForNewBalances() {
	if token == "" {
		fmt.Println("You need to log in first.")
		return
	}
	fmt.Println("Checking for new balances")

	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/assets?fields=balance,asset_type,user_id", apiURL), nil)
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
