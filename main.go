package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Owner struct {
	Login string
}

type Item struct {
	ID              int
	Name            string
	FullName        string `json:"full_name"`
	Owner           Owner
	Description     string
	CreatedAt       string `json:"created_at"`
	StargazersCount int    `json:"stargazers_count"`
}

func main() {
	var username string
	fmt.Print("Enter GitHub username: ")
	fmt.Scanln(&username)

	res, err := http.Get("https://api.github.com/users/" + username + "/repos")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("Unexpected status code: %d\nBody: %s", res.StatusCode, body)
	}

	var items []Item
	err = json.Unmarshal(body, &items)
	if err != nil {
		log.Fatal(err)
	}

	if len(items) == 0 {
		fmt.Println("No repositories found for user:", username)
		return
	}

	PrintResponse(items)
}
