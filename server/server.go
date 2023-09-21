package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func doWork(ctx context.Context, resChan chan int) {
	log.Println("[doWork] launch the doWork")
	sum := 0
	for {
		log.Println("[doWork] one iteration")
		time.Sleep(time.Millisecond)
		select {
		case <-ctx.Done():
			log.Println("[doWork] ctx Done is received inside doWork")
			return
		default:
			sum++
			if sum > 1000 {
				log.Println("[doWork] sum has reached 1000")
				resChan <- sum
				return
			}
		}
	}

}

func main() {
	log.SetOutput(os.Stdout)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("[handler] request received...")
		rCtx := r.Context()
		resChan := make(chan int)

		go doWork(rCtx, resChan)
		select {
		case <-rCtx.Done():
			log.Println("[handler] context canceled in main handler, client is disconnected")
		case result := <-resChan:
			log.Println("[handler] received 1000")
			log.Println("[handler] send response")
			fmt.Fprintf(w, "response %d", result)
			return
		}
	})

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
