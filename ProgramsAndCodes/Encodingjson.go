package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type student struct {
	Name   string `json:"name"`
	Course string `json:"course"`
}

func main() {
	stu := []student{
		{"Nitin", "MCA"},
		{"Shivank", "B-TECh"},
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(stu)
	f, err := os.Create("studentdb.json")

	if err != nil {
		panic(err)
	}
	defer f.Close()
	io.Copy(f, buf)

	fmt.Println(stu)

}
