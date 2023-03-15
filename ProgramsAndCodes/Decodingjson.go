package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type student struct {
	Name   string `json:"name"`
	Course string `json:"course"`
}

func main() {

	stu := []student{}

	f, err := os.Open("studentdb.json")
	json.NewDecoder(f).Decode(&stu)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Println(stu)
}
