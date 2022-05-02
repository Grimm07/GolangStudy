package Queue

import (
	"fmt"
	"math/rand"
	"time"
)

func addToQueue(queue[] string, person string) []string {
	queue = append(queue,person)
	fmt.Println("Person added to queue: ", person)
	return queue
}

func removeFromQueue(queue[] string) ([]string) {
	person := queue[0]
	fmt.Println("Person removed from queue: ",person)
	return queue[1:]
}

func Queue() {
	var queue[] string
	queue = addToQueue(queue,"Bob")
	queue = addToQueue(queue,"Carol")
	queue = addToQueue(queue,"George")
	queue = addToQueue(queue,"Judith")

	fmt.Println("\nWelcome to the waiting list!\n")
	fmt.Println("You have been added into the queue at number 5\nCurrent queue: ",queue)
	queue = addToQueue(queue,"You")
	fmt.Println("Updated queue: ",queue)
	for i := range queue {
		sleepTime := rand.Intn(5)
		time.Sleep(time.Duration(sleepTime)*time.Second)
		queue = removeFromQueue(queue)
		fmt.Println("Updated queue: ",queue)
		if i == 3 {
			fmt.Println("You're up!")
		}
	}
	fmt.Println("\nThanks for coming!")
}