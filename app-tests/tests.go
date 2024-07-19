package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL = "http://go-api:8080"

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Asset struct {
	AssetType string  `json:"assetType"`
	Balance   float64 `json:"balance"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

func createUser(user User) error {
	userJSON, _ := json.Marshal(user)
	resp, err := http.Post(baseURL+"/users", "application/json", bytes.NewBuffer(userJSON))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println("User created:", user.Username)
	return nil
}

func loginUser(user User) (string, error) {
	userJSON, _ := json.Marshal(user)
	resp, err := http.Post(baseURL+"/login", "application/json", bytes.NewBuffer(userJSON))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var tokenResp TokenResponse
	err = json.Unmarshal(body, &tokenResp)
	if err != nil {
		return "", err
	}
	return tokenResp.Token, nil
}

func createAsset(token string, asset Asset) (string, error) {
	assetJSON, _ := json.Marshal(asset)
	req, _ := http.NewRequest("POST", baseURL+"/assets", bytes.NewBuffer(assetJSON))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func getUserAssets(token string) (string, error) {
	req, _ := http.NewRequest("GET", baseURL+"/assets?fields=balance,asset_type", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func main() {
	users := map[string]User{
		"user1":  {Username: "user1", Password: "password1"},
		"user2":  {Username: "user2", Password: "password2"},
		"user3":  {Username: "user3", Password: "password3"},
		"user4":  {Username: "user4", Password: "password4"},
		"user5":  {Username: "user5", Password: "password5"},
		"user6":  {Username: "user6", Password: "password6"},
		"user7":  {Username: "user7", Password: "password7"},
		"user8":  {Username: "user8", Password: "password8"},
		"user9":  {Username: "user9", Password: "password9"},
		"user10": {Username: "user10", Password: "password10"},
	}

	orders := map[string][]Asset{
		"user1":  {{AssetType: "EUR", Balance: 1000}, {AssetType: "USD", Balance: 500}},
		"user2":  {{AssetType: "EUR", Balance: 2000}, {AssetType: "USD", Balance: 1500}},
		"user3":  {{AssetType: "EUR", Balance: 3000}, {AssetType: "USD", Balance: 2500}},
		"user4":  {{AssetType: "EUR", Balance: 4000}, {AssetType: "USD", Balance: 3500}},
		"user5":  {{AssetType: "EUR", Balance: 5000}, {AssetType: "USD", Balance: 4500}},
		"user6":  {{AssetType: "EUR", Balance: 6000}, {AssetType: "USD", Balance: 5500}},
		"user7":  {{AssetType: "EUR", Balance: 7000}, {AssetType: "USD", Balance: 6500}},
		"user8":  {{AssetType: "EUR", Balance: 8000}, {AssetType: "USD", Balance: 7500}},
		"user9":  {{AssetType: "EUR", Balance: 9000}, {AssetType: "USD", Balance: 8500}},
		"user10": {{AssetType: "EUR", Balance: 10000}, {AssetType: "USD", Balance: 9500}},
	}

	for username, user := range users {
		err := createUser(user)
		if err != nil {
			fmt.Println("Error creating user:", username, err)
			continue
		}

		token, err := loginUser(user)
		if err != nil {
			fmt.Println("Error logging in user:", username, err)
			continue
		}

		for _, order := range orders[username] {
			response, err := createAsset(token, order)
			if err != nil {
				fmt.Println("Error creating asset for user:", username, err)
				continue
			}
			fmt.Println("Asset creation response for user:", username, response)
		}

		assets, err := getUserAssets(token)
		if err != nil {
			fmt.Println("Error getting assets for user:", username, err)
			continue
		}
		fmt.Println("Assets for user:", username, assets)
	}
}
