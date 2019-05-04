package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getResponse(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error: ", err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error: ", err.Error())
		os.Exit(1)
	}
	return string(body)
}

func main() {
	response := getResponse("http://localhost:8000/api/book/1")
	log.Println(response)

}
