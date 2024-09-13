package main

import (
	"fmt"
	"log"
	"math/rand"
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

		numbers := []int{
			1,
			2,
			3,
			4,
			5,
			15,
		}

		messages := []string{
			"Example logging STDOUT",
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit...",
			"Beep Boop Beep Boop!",
			"Another example logging STDOUT",
			"Proin luctus quis velit nec congue.",
			"Phasellus ac ligula pulvinar, viverra odio ultricies, mattis sapien.",
		}

		for {
			fmt.Println("=================== TEST GO ROUTINE ===================")
			fmt.Println()
			fmt.Println("I can generate text just like a real AI too:")
			fmt.Println(messages[rand.Intn(len(messages))])
			fmt.Println(messages[rand.Intn(len(messages))])
			fmt.Println()

			fmt.Println("Performing critical business logics:")
			myNum := numbers[rand.Intn(len(numbers))]
			fmt.Println("Computing - what is: ", myNum)

			if myNum%3 == 0 && myNum%5 == 0 {
				fmt.Println(myNum, "is a FizzBuzz!")
			} else if myNum%3 == 0 {
				fmt.Println(myNum, "is a Fizz!")
			} else if myNum%5 == 0 {
				fmt.Println(myNum, "is a Buzz!")
			} else {
				fmt.Println(myNum, "is just a number :(")
			}

			fmt.Println()
			time.Sleep(second * seconds)
		}
	}()

	log.Println("Listening...")
	http.ListenAndServe(":8888", nil)
}
