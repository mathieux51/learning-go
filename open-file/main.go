package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// Users ...
type Users struct {
	Users []User `json:"users"`
}

// User ...
type User struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

func main() {
	jsonFile, err := os.Open("db.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users Users

	json.Unmarshal(byteValue, &users)

	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User ID: " + strconv.Itoa(users.Users[i].ID))
		fmt.Println("User Name: " + users.Users[i].Name)
	}

}
