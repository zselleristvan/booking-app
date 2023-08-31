package main

import "fmt"

func main() {

	var conferenceName = "Go conference"

	const conferenceTickets = 50

	var remainingTickets = 50
	fmt.Println("Welcome to", conferenceName, "booing application")
	fmt.Println("We have total of", conferenceTickets, "and tickets ", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")
}
