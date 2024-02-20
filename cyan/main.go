package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type appConfig struct {
	App      string                 `json:"app"`
	Entities map[string]interface{} `json:"entities"`
}

func main() {

	b, err := os.Open("cyan.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	var c appConfig
	json.NewDecoder(b).Decode(&c)

	fmt.Printf("Starting Cyan App: %s \n", c.App)

	for k, v := range c.Entities {
		fmt.Printf("key[%s] value[%s]\n", k, v)
	}

	// mux := http.NewServeMux()

	// fmt.Println("Hello, world.")
	// http.ListenAndServe(":3000", mux)
}
