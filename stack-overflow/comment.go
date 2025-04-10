package main

import (
	"sync"
	"time"
)

type Comment struct {
	Id        int
	Content   string
	Author    *User
	CreatedAt time.Time
	rw        sync.RWMutex
}

func NewComment(content string, author *User) *Comment {
	return &Comment{
		Id:        generateId(),
		Content:   content,
		Author:    author,
		CreatedAt: time.Now(),
	}
}

func (c *Comment) GetContent() string {
	c.rw.RLock()
	defer c.rw.RUnlock()
	return c.Content
}

func (c *Comment) GetAuthor() *User {
	c.rw.RLock()
	defer c.rw.RUnlock()
	return c.Author
}

func (c *Comment) GetCreatedAt() time.Time {
	c.rw.RLock()
	defer c.rw.RUnlock()
	return c.CreatedAt
}
