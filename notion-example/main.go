package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Parent struct {
	DatabaseId string `json:"database_id"`
}

type Properties struct {
	Title Title `json:"title"`
}

type Title struct {
	Title []RichText `json:"title"`
}

type RichText struct {
	Text Text `json:"text"`
}

type Text struct {
	Content string `json:"content"`
}

type Payload struct {
	Parent     Parent     `json:"parent"`
	Properties Properties `json:"properties"`
}

func main() {
	loadEnv()

	client := &http.Client{}
	url := "https://api.notion.com/v1/pages"

	payload := Payload{
		Parent: Parent{
			os.Getenv("DATABASE_ID"),
		},
		Properties: Properties{
			Title: Title{
				Title: []RichText{
					{
						Text: Text{
							Content: "Yurts in Big Sur, California",
						},
					},
				},
			},
		},
	}

	jsonBody, _ := json.Marshal(payload)

	requestBody := strings.NewReader(string(jsonBody))

	req, _ := http.NewRequest("POST", url, requestBody)

	req.Header.Add("Authorization", "Bearer "+os.Getenv("API_KEY"))
	req.Header.Add("Notion-Version", "2022-02-22")
	req.Header.Add("Content-Type", "application/json")

	res, _ := client.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}
}
