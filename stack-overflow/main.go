package main

import (
	"fmt"
)

func main() {

	so := NewStackOverflow()

	user1 := so.CreateUser("john_doe", "john@example.com")
	user2 := so.CreateUser("jane_smith", "jane@example.com")

	question := so.CreateQuestion(
		"How to implement a binary search tree in Go?",
		"I need help implementing a binary search tree in Go. Can someone provide an example?",
		user1,
		[]string{"go", "data-structures", "binary-search-tree"},
	)

	fmt.Printf("\nAfter creating question:\n")
	fmt.Printf("%s: %d\n", user1.GetUsername(), user1.GetReputation())
	fmt.Printf("%s: %d\n", user2.GetUsername(), user2.GetReputation())

	// Add a comment to the question
	so.CreateComment(
		"Have you tried looking at the standard library?",
		user2,
		question,
		nil,
	)

	fmt.Printf("\nAfter adding comment to question:\n")
	fmt.Printf("%s: %d\n", user1.GetUsername(), user1.GetReputation())
	fmt.Printf("%s: %d\n", user2.GetUsername(), user2.GetReputation())

	// Create an answer
	answer := so.CreateAnswer(
		"Here's a simple implementation of a binary search tree in Go...",
		user2,
		question,
	)

	fmt.Printf("\nAfter creating answer:\n")
	fmt.Printf("%s: %d\n", user1.GetUsername(), user1.GetReputation())
	fmt.Printf("%s: %d\n", user2.GetUsername(), user2.GetReputation())

	// Add a comment to the answer
	so.CreateComment(
		"Great answer! This helped me a lot.",
		user1,
		nil,
		answer,
	)

	fmt.Printf("\nAfter adding comment to answer:\n")
	fmt.Printf("%s: %d\n", user1.GetUsername(), user1.GetReputation())
	fmt.Printf("%s: %d\n", user2.GetUsername(), user2.GetReputation())

	// Vote on the question and answer using the system methods
	so.VoteOnQuestion(user2, question)
	so.VoteOnAnswer(user1, answer)

	fmt.Printf("\nAfter voting:\n")
	fmt.Printf("%s: %d\n", user1.GetUsername(), user1.GetReputation())
	fmt.Printf("%s: %d\n", user2.GetUsername(), user2.GetReputation())

	// Print the complete question details using the system method
	fmt.Println("\nQuestion Details:")
	so.PrintQuestionDetails(question.Id)

	// Print votes by user1
	fmt.Println("\nVotes by user1:")
	so.PrintUserVotes(user1.Id)

	// Print votes by user2
	fmt.Println("\nVotes by user2:")
	so.PrintUserVotes(user2.Id)

	// Search for questions
	searchResults := so.SearchQuestions("binary search tree")
	fmt.Println("\nSearch results for 'binary search tree':")
	for _, q := range searchResults {
		fmt.Printf("- %s\n", q.GetTitle())
	}

	// Search by tag
	tagResults := so.SearchByTag("go")
	fmt.Println("\nQuestions tagged with 'go':")
	for _, q := range tagResults {
		fmt.Printf("- %s\n", q.GetTitle())
	}
}
