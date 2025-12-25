package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HelloHandler called")
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	now := time.Now()

	fmt.Println("Current:", now.Format("2006-01-02 15:04:05"))

	oneHourLater := now.Add(1 * time.Hour)
	fmt.Println("One hour later:", oneHourLater.Format("2006-01-02 15:04:05"))

	oneHourAgo := now.Add(-1 * time.Hour)
	fmt.Println("One hour ago:", oneHourAgo.Format("2006-01-02 15:04:05"))

	diff := oneHourLater.Sub(oneHourAgo).Seconds()
	fmt.Println("Difference:", diff)

	start := time.Now()
	time.Sleep(100 * time.Millisecond)
	elpased := time.Since(start)
	fmt.Println("Elpased:", elpased.Seconds())

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Worker", i, "started")
			time.Sleep(1 * time.Second)
		}()
	}
	wg.Wait()
	fmt.Println("All workers finished")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	select {
	case <-time.After(4 * time.Second):
		fmt.Println("Timeout")
	case <-ctx.Done():
		fmt.Println("Context cancelled:", ctx.Err())
	}

	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server is running on port 8080, http://localhost:8080/hello?name=John")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
