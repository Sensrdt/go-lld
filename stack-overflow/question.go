package main

import (
	"sync"
	"time"
)

type Question struct {
	Id        int
	Title     string
	Content   string
	Author    *User
	Tags      []string
	Answers   []*Answer
	Comments  []*Comment
	Vote      *Vote
	CreatedAt time.Time
	rw        sync.RWMutex
}

func NewQuestion(title, content string, author *User, tags []string) *Question {
	question := &Question{
		Id:        generateId(),
		Title:     title,
		Content:   content,
		Author:    author,
		Tags:      tags,
		Answers:   make([]*Answer, 0),
		Comments:  make([]*Comment, 0),
		Vote:      NewVote(),
		CreatedAt: time.Now(),
	}
	author.AddQuestion(question)
	return question
}

var _ VoteInterface = (*Question)(nil)

// Implement VoteInterface
func (q *Question) AddVote(user *User) error {
	q.rw.Lock()
	defer q.rw.Unlock()
	return q.Vote.AddVoteCount(user.Id, Upvote)
}

func (q *Question) RemoveVote(user *User) error {
	q.rw.Lock()
	defer q.rw.Unlock()
	return q.Vote.RemoveVoteCount(user.Id)
}

func (q *Question) GetUserVote(user *User) (VoteType, bool) {
	q.rw.RLock()
	defer q.rw.RUnlock()
	return q.Vote.GetSpecificUserVote(user)
}

func (q *Question) GetVoteCount() int {
	q.rw.RLock()
	defer q.rw.RUnlock()
	return q.Vote.GetTotalVoteCount()
}

// Answer methods
func (q *Question) AddAnswer(answer *Answer) {
	q.rw.Lock()
	defer q.rw.Unlock()
	q.Answers = append(q.Answers, answer)
}

func (q *Question) GetAnswers() []*Answer {
	q.rw.RLock()
	defer q.rw.RUnlock()
	return q.Answers
}

// Comment methods
func (q *Question) AddComment(comment *Comment) {
	q.rw.Lock()
	defer q.rw.Unlock()
	q.Comments = append(q.Comments, comment)
	q.Author.AddComment(comment)
}

func (q *Question) GetComments() []*Comment {
	q.rw.RLock()
	defer q.rw.RUnlock()
	return q.Comments
}

// Tag methods
func (q *Question) AddTag(tag string) {
	q.rw.Lock()
	defer q.rw.Unlock()
	q.Tags = append(q.Tags, tag)
}

func (q *Question) GetTags() []string {
	q.rw.RLock()
	defer q.rw.RUnlock()
	return q.Tags
}

// Getters
func (q *Question) GetTitle() string {
	q.rw.RLock()
	defer q.rw.RUnlock()
	return q.Title
}

func (q *Question) GetContent() string {
	q.rw.RLock()
	defer q.rw.RUnlock()
	return q.Content
}

func (q *Question) GetAuthor() *User {
	q.rw.RLock()
	defer q.rw.RUnlock()
	return q.Author
}

func (q *Question) GetCreatedAt() time.Time {
	q.rw.RLock()
	defer q.rw.RUnlock()
	return q.CreatedAt
}
