package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Person struct {
	Name       string `json:"name,omitempty"`
	Age        int    `json:"age,omitempty"`
	Creditcard string `json:"-"`
}

func main() {
	p := []Person{
		{"Nitin", 24, "Axis credit"},
		{"Shivank", 24, "Pnb credit"},
	}
	//buf := new(bytes.Buffer)
	pbytes, err := json.Marshal(p)
	f, err1 := os.Create("person.json")
	if err != nil {
		panic(err)
	}
	if err1 != nil {
		panic(err1)
	}

	defer f.Close()
	//fmt.Println(string(pbytes))
	// i := strings.NewReader(string(pbytes))
	//io.Copy(buf, bytes.NewBufferString(string(pbytes)))
	//fmt.Println(buf)
	io.Copy(f, strings.NewReader(string(pbytes)))

	p2 := []Person{}

	f2, err2 := os.Open("person.json")
	if err2 != nil {
		panic(err2)
	}

	var buff strings.Builder
	_, err4 := io.Copy(&buff, f2)

	if err4 != nil {
		panic(err4)
	}

	err3 := json.Unmarshal([]byte(buff.String()), &p2)

	log.Print(err3)

	fmt.Println(p2)
}
