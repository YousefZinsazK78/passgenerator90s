package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)
	req, err := http.NewRequest("GET", "http://localhost:8000/", nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	log.Println(resp)
}
