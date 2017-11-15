package main

import (
	"regexp"
	"time"
	"fmt"
	"math/rand"
)

func ElizaResponse(input string) string {
	//array of responses the eliza program will respond with
	responses := []string{
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}

	//Adapted from https://golang.org/pkg/regexp/
	//Searchs the input to look for the word "Father" on its own and assigns it to this variable
	father, _ := regexp.MatchString("(?i)\\bfather\\b", input)
	
	//if the user input contains the word "father" it will return this string
	if (father) {
		return("Why don’t you tell me more about your father?")
	}

	//returns the responses to the main function
	return responses[rand.Intn(len(responses))]
}


func main() {
	//To pick a random response from the ElizaResponse func
	rand.Seed(time.Now().UTC().UnixNano())

	//Prints each input to the console and passes it into the ElizaResponse func
	fmt.Println("\nInput: " + "People say I look like both my mother and father.")
	fmt.Println("Output: " + ElizaResponse("People say I look like both my mother and father."))

	fmt.Println("\nInput: " + "Father was a teacher.")
	fmt.Println("Output: " + ElizaResponse("Father was a teacher."))

	fmt.Println("\nInput: " + "I was my father’s favourite.")
	fmt.Println("Output: " + ElizaResponse("I was my father’s favourite."))

	fmt.Println("\nInput: " + "I'm looking forward to the weekend.")
	fmt.Println("Output: " + ElizaResponse("I'm looking forward to the weekend."))

	fmt.Println("\nInput: " + "My grandfather was French!")
	fmt.Println("Output: " + ElizaResponse("My grandfather was French!"))
}