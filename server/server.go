package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	log.SetOutput(os.Stdout)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("request received...")
		time.Sleep(time.Minute * 1)
		fmt.Fprint(w, "custom response")
		log.Println("response sent...")
	})

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
