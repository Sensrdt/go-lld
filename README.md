# Go Low-Level Design (LLD) & Interview Prep 🚀

This repository contains various Low-Level Design (LLD) implementations, concurrency patterns, and system design problems written in Go. It serves as a great practice ground for Go backend interviews.

Here is a quick overview of what is implemented in each folder:

## 📐 System Design & LLD
- **`atm/`**: Low-Level Design of an ATM Machine. Likely utilizes the **State Design Pattern** to manage states like `HasCard`, `SelectingTransaction`, and `DispensingCash`.
- **`parking-lot/`**: The classic Parking Lot LLD. Handles ticket generation, multiple vehicle sizes, spot allocation, and payment calculations.
- **`stack-overflow/`**: LLD for a Stack Overflow clone. Manages entities like `Users`, `Questions`, `Answers`, `Comments`, and `Bounties/Voting`.
- **`vending-machine/`**: LLD of a Vending Machine using the State pattern to handle money insertion, item selection, and dispensing.

## ⚡ Concurrency & Network Patterns
- **`http-failover/`**: Implementation of **Hedged Requests (Dynamic Failover)**. Sends staggered HTTP requests using `time.Timer` and `context.Context` to fetch from multiple URLs, immediately returning the fastest successful response while cleanly canceling the rest.
- **`rate-limiter/`**: A simple interval-based request rate limiter (e.g., max 3 requests every 2 seconds) using channels and timers.
- **`tokenBucketRatelimiter.go/`**: A more advanced Rate Limiter implementation utilizing the **Token Bucket Algorithm** (allowing bursts up to a capacity, refilling at a steady rate).
- **`threadSafeQueue.go/`**: Implementation of a thread-safe Queue data structure, ensuring safe concurrent access using `sync.Mutex` or Go Channels.
- **`concurrency/`**: A collection of common Go concurrency patterns, sync primitives, and exercises.

## 🛠️ Core Go Concepts
- **`interfaces/`**: Demonstrations of Go's powerful interface system. Includes decoupled designs like a `NotificationService` (Email/SMS) and a clean `InMemCache`.
- **`oops/`**: Demonstrates how to apply Object-Oriented Programming (OOP) concepts—like Abstraction, Encapsulation, and Polymorphism—using Go structs and interfaces.
- **`design-patterns/`**: Common Gang of Four (GoF) design patterns implemented cleanly in Go (e.g., Singleton, Factory, Builder, Observer).
- **`rest-api/`**: Basic implementation patterns for structuring and building RESTful HTTP APIs in Go.

## 🧩 Interview Practice & Brain Teasers
- **`trivial/`**: Small, tricky interview questions. For example, printing the alphabet asynchronously using single/multiple goroutines, daisy-chained channels, and `sync.WaitGroup` traps.
- **`interview/`**: Miscellaneous scripts and logic exercises targeted specifically at cracking Go coding interviews.