package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan<-true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(2 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true
	close(doneChan) //close the channel in the op. that takes the most time
}

func main() {
	//dones := make([] chan bool, 4)
	isDone:=make(chan bool)
	// dones[0] = make(chan bool)
	// dones[1] = make(chan bool)
	// dones[2] = make(chan bool)
	// dones[3] = make(chan bool)
	go greet("Nice to meet you!", isDone)
	go greet("How are you?", isDone)
	go slowGreet("⌛ How ... are ... you ...?", isDone)
	go greet("I'm liking the GoLang!", isDone)
	//fmt.Println(<- isDone) //Optnl.

	// for _, done := range dones{
	// 	<-done
	// }

	for range isDone{
		//fmt.Println(doneChan)
	}

}
