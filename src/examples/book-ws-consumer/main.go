package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
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

func getResponseWithSettings(url string) string {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    1 * time.Millisecond,
		DisableCompression: true,
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
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
	response1 := getResponse("http://localhost:8000/api/book/1")
	log.Println("Response ==> ",response1)

	response2 := getResponseWithSettings("http://localhost:8000/api/book/2")
	log.Println("Response ==> ",response2)

}
