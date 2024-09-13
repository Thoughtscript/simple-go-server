package main

import (
	"fmt"
	"log"
	"net/http"
	h "simplegoserver/handlers"
	"time"
)

func main() {
	fileServer := http.FileServer(http.Dir("./assets"))
	http.Handle("/", http.StripPrefix("/public/", fileServer))
	http.Handle("/public/", http.StripPrefix("/public/", fileServer))

	http.HandleFunc("/heartbeat", h.HeartBeat)
	http.HandleFunc("/stop", h.ShutDown)

	go func() {
		var second time.Duration = 1000000000
		var seconds time.Duration = 4

		for {
			fmt.Println("=================== TEST GO ROUTINE ===================")
			fmt.Println("Example logging STDOUT")
			fmt.Println("Lorem ipsum dolor sit amet, consectetur adipiscing elit...")
			fmt.Println("Fizz")
			fmt.Println("Buzz")
			fmt.Println("Fizz Buzz")
			fmt.Println("Example logging STDOUT")
			time.Sleep(second * seconds)
		}
	}()

	log.Println("Listening...")
	http.ListenAndServe(":8888", nil)
}
