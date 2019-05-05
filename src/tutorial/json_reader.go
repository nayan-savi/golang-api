package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

var users Users

func main() {
	jsonFile, err := os.Open("users.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	json.Unmarshal(byteValue, &users)

	for i := 0; i < len(users.Users); i++ {
		fmt.Println("Name: " + users.Users[i].Name)
		fmt.Println("Age: " + strconv.Itoa(users.Users[i].Age))
		fmt.Println("Type: " + users.Users[i].Type)
		fmt.Println("Facebook: " + users.Users[i].Social.Facebook)
		fmt.Println("Twitter: " + users.Users[i].Social.Twitter)
	}
	fmt.Println(result["users"])

}
