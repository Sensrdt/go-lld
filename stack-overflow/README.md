# Stack Overflow Low-Level Design (LLD)

This folder implements a mock backend Low-Level Design for a Stack Overflow-style Q&A platform.

## File Implementations

- **`system.go`**: Defines the central `StackOverflow` struct, storing all data in memory and protecting access with a `sync.RWMutex`.
- **`service.go`**: Contains the core business logic functions for interacting with the system.
- **`user.go`**: Defines the `User` struct and tracks user reputation and stats.
- **`question.go`**: Defines the `Question` struct, linking it to answers, comments, tags, and authors.
- **`answer.go`**: Defines the `Answer` struct, linking it to its parent question and comments.
- **`comment.go`**: Defines the `Comment` struct for nested discussions.
- **`vote.go`**: Defines `VoteType` and the `VoteInterface` for handling upvotes/downvotes across posts.
- **`tag.go`**: Defines the `Tag` struct for categorizing questions.
- **`main.go`**: Simulates the system by initializing users, asking a question, and adding answers.
