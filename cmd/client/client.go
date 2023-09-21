package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	log.SetOutput(os.Stdout)
	ctx := context.Background()

	req, err := http.NewRequest("GET", "http://localhost:8000/", nil)
	if err != nil {
		panic(err)
	}
	cancelCtx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()

	req = req.WithContext(cancelCtx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	log.Println(resp)
}
