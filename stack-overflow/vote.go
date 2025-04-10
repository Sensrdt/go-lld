package main

import "sync"

type VoteType int

const (
	Upvote VoteType = iota + 1
	Downvote
)

type VoteInterface interface {
	AddVote(user *User) error
	RemoveVote(user *User) error
	GetUserVote(user *User) (VoteType, bool)
	GetVoteCount() int
}

type Vote struct {
	Count  int
	Voters map[int]VoteType
	rw     sync.RWMutex
}

func NewVote() *Vote {
	return &Vote{
		Count:  0,
		Voters: make(map[int]VoteType),
	}
}

func (v *Vote) AddVoteCount(userId int, voteType VoteType) error {
	v.rw.Lock()
	defer v.rw.Unlock()

	if existingVote, exists := v.Voters[userId]; exists {
		if existingVote == voteType {
			return nil // Already voted the same way
		}
		// Remove the opposite vote
		if existingVote == Upvote {
			v.Count--
		} else {
			v.Count++
		}
	} else {
		if voteType == Upvote {
			v.Count++
		} else {
			v.Count--
		}
	}

	v.Voters[userId] = voteType
	return nil
}

func (v *Vote) RemoveVoteCount(userId int) error {
	v.rw.Lock()
	defer v.rw.Unlock()

	if voteType, exists := v.Voters[userId]; exists {
		if voteType == Upvote {
			v.Count--
		} else {
			v.Count++
		}
		delete(v.Voters, userId)
	}
	return nil
}

func (v *Vote) GetSpecificUserVote(user *User) (VoteType, bool) {
	v.rw.RLock()
	defer v.rw.RUnlock()
	voteType, exists := v.Voters[user.Id]
	return voteType, exists
}

func (v *Vote) GetTotalVoteCount() int {
	v.rw.RLock()
	defer v.rw.RUnlock()
	return v.Count
}
