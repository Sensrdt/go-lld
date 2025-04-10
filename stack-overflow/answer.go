package main

import (
	"sync"
	"time"
)

type Answer struct {
	Id        int
	Content   string
	Author    *User
	Question  *Question
	Comments  []*Comment
	Vote      *Vote
	CreatedAt time.Time
	rw        sync.RWMutex
}

func NewAnswer(content string, author *User, question *Question) *Answer {
	answer := &Answer{
		Id:        generateId(),
		Content:   content,
		Author:    author,
		Question:  question,
		Comments:  make([]*Comment, 0),
		Vote:      NewVote(),
		CreatedAt: time.Now(),
	}
	author.AddAnswer(answer)
	return answer
}

var _ VoteInterface = (*Answer)(nil)

// Implement VoteInterface
func (a *Answer) AddVote(user *User) error {
	a.rw.Lock()
	defer a.rw.Unlock()
	return a.Vote.AddVoteCount(user.Id, Upvote)
}

func (a *Answer) RemoveVote(user *User) error {
	a.rw.Lock()
	defer a.rw.Unlock()
	return a.Vote.RemoveVoteCount(user.Id)
}

func (a *Answer) GetUserVote(user *User) (VoteType, bool) {
	a.rw.RLock()
	defer a.rw.RUnlock()
	return a.Vote.GetSpecificUserVote(user)
}

func (a *Answer) GetVoteCount() int {
	a.rw.RLock()
	defer a.rw.RUnlock()
	return a.Vote.GetTotalVoteCount()
}

// Comment methods
func (a *Answer) AddComment(comment *Comment) {
	a.rw.Lock()
	defer a.rw.Unlock()
	a.Comments = append(a.Comments, comment)
	a.Author.AddComment(comment)
}

func (a *Answer) GetComments() []*Comment {
	a.rw.RLock()
	defer a.rw.RUnlock()
	return a.Comments
}

// Getters
func (a *Answer) GetContent() string {
	a.rw.RLock()
	defer a.rw.RUnlock()
	return a.Content
}

func (a *Answer) GetAuthor() *User {
	a.rw.RLock()
	defer a.rw.RUnlock()
	return a.Author
}

func (a *Answer) GetCreatedAt() time.Time {
	a.rw.RLock()
	defer a.rw.RUnlock()
	return a.CreatedAt
}
