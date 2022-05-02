/*
***********************************************************************************************************************
*	Go language study main testing file																				  *
*	This file contains functions to test the programs outlined in the 327ProjectPlan.docx							  *
*																													  *
*	NOTE: To add packages to this file, run go install %PATH_TO_PACKAGE% - Provide it with an absolute path		      *
**********************************************************************************************************************
 */

package main

import (
	Client_Server "Programs/Client-Server"
	"Programs/Hello-World"
	Queue "Programs/Queue"
	"fmt"
	"strings"
)

// mapping for print statements (G == SUCCESS, R == FAILURE, W == NONE)

type RetVal int

// Return Values - used to denote if a function passes the test or not
const (
	SUCCESS RetVal = iota // G
	FAILURE               // R
	NONE                  // W
)

// here we associate the above constants with their respective colors and output the messages
func printTest(message string) {
	var pColor string
	var rVal RetVal
	if strings.Contains(message, "FAILURE") {
		rVal = FAILURE
	} else if strings.Contains(message, "SUCCESS") {
		rVal = SUCCESS
	} else {
		rVal = NONE
	}

	switch rVal {
	case FAILURE:
		pColor = string("\033[31m") // R (FAILURE)
		fmt.Println(pColor + message)
		break
	case SUCCESS:
		pColor = string("\033[32m") // G (SUCCESS)
		fmt.Println(pColor + message)
		break
	default:
		pColor = string("\033[0m") // W (NONE)
		fmt.Println(pColor + message)
		break
	}
}

// Function Tests a simple "hello world!" program (Should never fail)
func hello() {
	msg := Hello.Hello("Trystan")
	var test string
	if strings.Contains(msg, "Hello") {
		test = "SUCCESS: HELLO FUNCTION PASSED TEST"
	} else {
		test = "FAILURE: HELLO FUNCTION FAILED TEST"
	}
	printTest(test)

}

func queueCall() {
	Queue.Queue()
}

func ServerClient() {
	s := Client_Server.Server{}

	// start new coroutine to initialize server
	go s.StartServer()
	var clients [10]*Client_Server.Client

	for i := 0; i < 10; i++ {
		clients[i] = Client_Server.StartClient(i)
	}
	for i := 0; i < 10; i++ {
		// all of these messages should be logged
		testString := fmt.Sprintf("This is a test message from %d client.\n", i)
		clients[i].SendMessage(testString)
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("FROM SERVER: " + clients[i].ReceiveMessage() + "\n")
	}
	println("Closing Clients \n")
	for i := 0; i < 10; i++ {
		clients[i].CloseClient()
	}
	s.CloseServer()

}

// test method
func main() {
	//hello()
	queueCall()
	//ServerClient()
}
