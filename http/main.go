package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := http.Client{}
	resp, err := client.Get("https://text.npr.org/")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		bodyString := string(bodyBytes)
		log.Println(bodyString)
	}
}
