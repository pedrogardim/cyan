package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-cyan/cyan/db"
	"github.com/go-cyan/cyan/handler"
	"github.com/go-cyan/cyan/repository"
)

type appConfig struct {
	App   string `json:"app"`
	Mongo struct {
		Uri string `json:"uri"`
		Db  string `json:"db"`
	} `json:"mongo"`
	Entities map[string]interface{} `json:"entities"`
}

func pathToName(inputPath string) string {
	baseName := strings.TrimSuffix(filepath.Base(inputPath), filepath.Ext(inputPath))
	return baseName
}

func main() {

	b, err := os.Open("cyan.json")

	if err != nil {
		fmt.Println(err)
	}

	var c appConfig
	json.NewDecoder(b).Decode(&c)

	mongoClient := db.NewConnection(c.Mongo.Uri)

	fmt.Printf("Starting Cyan App: %s \n", c.App)

	mux := http.NewServeMux()

	for k, v := range c.Entities {
		name := pathToName(k)
		repo := repository.NewRepository(name, v, mongoClient, c.Mongo.Db)
		handler := handler.NewHandler(name, repo)
		mux.Handle("api/v1/"+name, handler)
	}

	fmt.Println("Hello, world.")
	http.ListenAndServe(":3000", mux)
}
