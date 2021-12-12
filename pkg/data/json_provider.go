package data

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func post(query string) {
	postBody, _ := json.Marshal(query)
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("https://api.spacex.land/graphql/", "application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	log.Printf(sb)
}
