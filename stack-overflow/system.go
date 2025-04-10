package main

import (
	"fmt"
	"strings"
	"sync"
)

type StackOverflow struct {
	users     map[int]*User
	questions map[int]*Question
	answers   map[int]*Answer
	comments  map[int]*Comment
	rw        sync.RWMutex
}

func NewStackOverflow() *StackOverflow {
	return &StackOverflow{
		users:     make(map[int]*User),
		questions: make(map[int]*Question),
		answers:   make(map[int]*Answer),
		comments:  make(map[int]*Comment),
	}
}

// User management
func (s *StackOverflow) CreateUser(username, email string) *User {
	s.rw.Lock()
	defer s.rw.Unlock()
	user := NewUser(username, email)
	s.users[user.Id] = user
	return user
}

func (s *StackOverflow) GetUser(id int) (*User, bool) {
	s.rw.RLock()
	defer s.rw.RUnlock()
	user, exists := s.users[id]
	return user, exists
}

// Question management
func (s *StackOverflow) CreateQuestion(title, content string, author *User, tags []string) *Question {
	s.rw.Lock()
	defer s.rw.Unlock()
	question := NewQuestion(title, content, author, tags)
	s.questions[question.Id] = question
	return question
}

func (s *StackOverflow) GetQuestion(id int) (*Question, bool) {
	s.rw.RLock()
	defer s.rw.RUnlock()
	question, exists := s.questions[id]
	return question, exists
}

// Answer management
func (s *StackOverflow) CreateAnswer(content string, author *User, question *Question) *Answer {
	s.rw.Lock()
	defer s.rw.Unlock()
	answer := NewAnswer(content, author, question)
	s.answers[answer.Id] = answer
	question.AddAnswer(answer)

	// Update reputation
	author.AddReputation(5) // Answer author gets reputation for answering

	return answer
}

func (s *StackOverflow) GetAnswer(id int) (*Answer, bool) {
	s.rw.RLock()
	defer s.rw.RUnlock()
	answer, exists := s.answers[id]
	return answer, exists
}

// Comment management
func (s *StackOverflow) CreateComment(content string, author *User, question *Question, answer *Answer) *Comment {
	s.rw.Lock()
	defer s.rw.Unlock()
	comment := NewComment(content, author)
	s.comments[comment.Id] = comment

	if question != nil {
		question.AddComment(comment)
		// Question author gets reputation for receiving a comment
		question.GetAuthor().AddReputation(2)
	} else if answer != nil {
		answer.AddComment(comment)
		// Answer author gets reputation for receiving a comment
		answer.GetAuthor().AddReputation(2)
	}

	// Comment author gets reputation for adding a comment
	author.AddReputation(1)

	return comment
}

func (s *StackOverflow) GetComment(id int) (*Comment, bool) {
	s.rw.RLock()
	defer s.rw.RUnlock()
	comment, exists := s.comments[id]
	return comment, exists
}

// Voting management
func (s *StackOverflow) VoteOnQuestion(user *User, question *Question) error {
	s.rw.Lock()
	defer s.rw.Unlock()

	err := question.AddVote(user)
	if err == nil {
		// Update reputation
		question.GetAuthor().AddReputation(10) // Question author gets reputation for upvote
		user.AddReputation(1)                  // Voter gets reputation for voting

		// Record the vote
		user.AddVote(question.Id, 0, Upvote)
	}
	return err
}

func (s *StackOverflow) VoteOnAnswer(user *User, answer *Answer) error {
	s.rw.Lock()
	defer s.rw.Unlock()

	err := answer.AddVote(user)
	if err == nil {
		// Update reputation
		answer.GetAuthor().AddReputation(10) // Answer author gets reputation for upvote
		user.AddReputation(1)                // Voter gets reputation for voting

		// Record the vote
		user.AddVote(0, answer.Id, Upvote)
	}
	return err
}

// Print user votes
func (s *StackOverflow) PrintUserVotes(userId int) {
	s.rw.RLock()
	defer s.rw.RUnlock()

	user, exists := s.users[userId]
	if !exists {
		fmt.Println("User not found")
		return
	}

	votes := user.GetVotes()

	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Printf("Votes by %s (ID: %d):\n", user.GetUsername(), user.Id)
	fmt.Println(strings.Repeat("-", 80))

	if len(votes) == 0 {
		fmt.Println("No votes cast yet")
	} else {
		for _, vote := range votes {
			if vote.QuestionId > 0 {
				question, exists := s.questions[vote.QuestionId]
				if exists {
					fmt.Printf("Question: %s (ID: %d) - Vote: %d\n",
						question.GetTitle(),
						question.Id,
						vote.VoteType)
				}
			} else if vote.AnswerId > 0 {
				answer, exists := s.answers[vote.AnswerId]
				if exists {
					question, _ := s.questions[answer.Question.Id]
					fmt.Printf("Answer to '%s' (Question ID: %d) - Vote: %d\n",
						question.GetTitle(),
						question.Id,
						vote.VoteType)
				}
			}
		}
	}

	fmt.Println(strings.Repeat("=", 80))
}

// Search functionality
func (s *StackOverflow) SearchQuestions(keyword string) []*Question {
	s.rw.RLock()
	defer s.rw.RUnlock()

	var results []*Question
	for _, question := range s.questions {
		if contains(question.Title, keyword) || contains(question.Content, keyword) {
			results = append(results, question)
		}
	}
	return results
}

func (s *StackOverflow) SearchByTag(tag string) []*Question {
	s.rw.RLock()
	defer s.rw.RUnlock()

	var results []*Question
	for _, question := range s.questions {
		for _, t := range question.Tags {
			if t == tag {
				results = append(results, question)
				break
			}
		}
	}
	return results
}

// Print question details
func (s *StackOverflow) PrintQuestionDetails(questionId int) {
	s.rw.RLock()
	defer s.rw.RUnlock()

	question, exists := s.questions[questionId]
	if !exists {
		fmt.Println("Question not found")
		return
	}

	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Printf("Question: %s\n", question.Title)
	fmt.Println(strings.Repeat("-", 80))

	fmt.Printf("Content: %s\n\n", question.Content)

	fmt.Printf("Author: %s (ID: %d)\n", question.Author.GetUsername(), question.Author.Id)
	fmt.Printf("Created At: %s\n", question.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Printf("Votes: %d\n\n", question.GetVoteCount())

	fmt.Println("Tags:")
	for _, tag := range question.Tags {
		fmt.Printf("  #%s\n", tag)
	}

	fmt.Println("\nComments:")
	if len(question.Comments) == 0 {
		fmt.Println("  No comments yet")
	} else {
		for _, comment := range question.Comments {
			fmt.Printf("  - %s (by %s): %s\n",
				comment.GetCreatedAt().Format("2006-01-02 15:04:05"),
				comment.GetAuthor().GetUsername(),
				comment.GetContent())
		}
	}

	fmt.Println("\nAnswers:")
	if len(question.Answers) == 0 {
		fmt.Println("  No answers yet")
	} else {
		for _, answer := range question.Answers {
			fmt.Printf("  - By %s (ID: %d)\n", answer.GetAuthor().GetUsername(), answer.GetAuthor().Id)
			fmt.Printf("    Content: %s\n", answer.GetContent())
			fmt.Printf("    Votes: %d\n", answer.GetVoteCount())

			fmt.Println("    Comments:")
			for _, comment := range answer.GetComments() {
				fmt.Printf("      - %s (by %s): %s\n",
					comment.GetCreatedAt().Format("2006-01-02 15:04:05"),
					comment.GetAuthor().GetUsername(),
					comment.GetContent())
			}
			fmt.Println()
		}
	}

	fmt.Println(strings.Repeat("=", 80))
}

// Helper function
func contains(s, substr string) bool {
	return len(substr) == 0 || len(s) >= len(substr) && s[0:len(substr)] == substr
}
