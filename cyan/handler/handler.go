package handler

import (
	"fmt"
	"net/http"
)

func CreateHandler(x interface{}) {
	mux := http.NewServeMux()

	fmt.Println("Hello, world.")
	http.ListenAndServe(":3000", mux)
}
