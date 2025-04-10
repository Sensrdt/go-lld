package main

import (
	"sync"
)

/*
	User
		Can ask questions
		Can answer questions
		Can vote on questions and answers
		Can add tags to questions
		Can add comments to questions and answers
		Can add links to questions and answers
		Can add images to questions and answers
*/

type VoteRecord struct {
	QuestionId int
	AnswerId   int
	VoteType   VoteType
}

type User struct {
	Id         int
	Username   string
	Email      string
	Reputation int
	Questions  []*Question
	Answers    []*Answer
	Comments   []*Comment
	Votes      []VoteRecord
	rw         sync.RWMutex
}

func NewUser(username, email string) *User {
	return &User{
		Id:         generateId(),
		Username:   username,
		Email:      email,
		Reputation: 1, // Starting reputation
		Questions:  make([]*Question, 0),
		Answers:    make([]*Answer, 0),
		Comments:   make([]*Comment, 0),
		Votes:      make([]VoteRecord, 0),
	}
}

func (u *User) AddReputation(points int) {
	u.rw.Lock()
	defer u.rw.Unlock()
	u.Reputation += points
}

func (u *User) GetReputation() int {
	u.rw.RLock()
	defer u.rw.RUnlock()
	return u.Reputation
}

func (u *User) AddQuestion(q *Question) {
	u.rw.Lock()
	defer u.rw.Unlock()
	u.Questions = append(u.Questions, q)
}

func (u *User) AddAnswer(a *Answer) {
	u.rw.Lock()
	defer u.rw.Unlock()
	u.Answers = append(u.Answers, a)
}

func (u *User) AddComment(c *Comment) {
	u.rw.Lock()
	defer u.rw.Unlock()
	u.Comments = append(u.Comments, c)
}

func (u *User) AddVote(questionId, answerId int, voteType VoteType) {
	u.rw.Lock()
	defer u.rw.Unlock()
	u.Votes = append(u.Votes, VoteRecord{
		QuestionId: questionId,
		AnswerId:   answerId,
		VoteType:   voteType,
	})
}

func (u *User) GetVotes() []VoteRecord {
	u.rw.RLock()
	defer u.rw.RUnlock()
	return u.Votes
}

// Getters
func (u *User) GetUsername() string {
	u.rw.RLock()
	defer u.rw.RUnlock()
	return u.Username
}

func (u *User) GetQuestions() []*Question {
	u.rw.RLock()
	defer u.rw.RUnlock()
	return u.Questions
}

func (u *User) GetAnswers() []*Answer {
	u.rw.RLock()
	defer u.rw.RUnlock()
	return u.Answers
}

func (u *User) GetComments() []*Comment {
	u.rw.RLock()
	defer u.rw.RUnlock()
	return u.Comments
}

// Helper function to generate unique IDs
var currentId = 0
var idMutex sync.Mutex

func generateId() int {
	idMutex.Lock()
	defer idMutex.Unlock()
	currentId++
	return currentId
}
