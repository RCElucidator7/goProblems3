package main

import (
	"regexp"
	"time"
	"fmt"
	"math/rand"
	"strings"
)

func ElizaResponse(input string) string {
	//array of responses the eliza program will respond with
	responses := []string{
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}

	iam := regexp.MustCompile("(?i)i(?:'| a|)?m(.*)")
	if iam.MatchString(input) {
		return iam.ReplaceAllString(input, "How do I know you are $1")
	}

	//Adapted from https://golang.org/pkg/regexp/
	//Searchs the input to look for the word "Father" on its own and assigns it to this variable
	father, _ := regexp.MatchString(`(?i)\\bfather\\b`, input)
	
	//if the user input contains the word "father" it will return this string
	if (father) {
		return("Why don’t you tell me more about your father?")
	}

	//returns the responses to the main function
	return responses[rand.Intn(len(responses))]
}

func reflection(input string) string{

	// List the pronouns to switch
	pronouns := [][]string{
		{`I`, `you`},
		{`me`, `you`},
		{`you`, `me`},
		{`my`, `your`},
		{`your`, `my`},
	}

	// Split input into values
	boundaries := regexp.MustCompile(`\b`)

	values := boundaries.Split(input, -1)

	//Loop through the range of values and reflect the pronoun if it finds a match
	for i, token := range values {
		for _, reflection := range pronouns {
			if matched, _ := regexp.MatchString(reflection[0], token); matched {
				values[i] = reflection[1]
				break
			}
		}
	}
	
	//Join the string of values back together
	return strings.Join(values, ``)
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

	fmt.Println("\nInput: " + "I am happy.")
	fmt.Println("Output: " + ElizaResponse("I am happy."))

	fmt.Println("\nInput: " + "I'm not happy with your responses")
	fmt.Println("Output: " + ElizaResponse("I'm not happy with your responses"))

	fmt.Println("\nInput: " + "“I AM not sure that you understand the effect that your questions are having on me.”")
	fmt.Println("Output: " + ElizaResponse("I AM not sure that you understand the effect that your questions are having on me."))

	fmt.Println("\nInput: " + "Im supposed to just take what you’re saying at face value?")
	fmt.Println("Output: " + ElizaResponse("Im supposed to just take what you’re saying at face value?"))

	fmt.Println(reflection("You are my friend cause I like you."))
}