package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("supersecretkey#123")

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// this will send the token to the client
func dummyLoginHandler(r *gin.Context) {
	token, _ := createToken()
	r.JSON(http.StatusCreated, gin.H{"token": token})
}
func createToken() (string, error) {
	exp := time.Now().Add(24 * time.Hour)
	claim := Claims{
		Username: "sridip",
		Role:     "admin",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	return token.SignedString(jwtKey)
}

// request logger
func requestLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Printf("Incoming request: %s - %s", ctx.Request.Method, ctx.Request.URL.Path)
		ctx.Next()
	}
}

// sample token validator middleware
func tokenValidatorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			log.Fatal("error: %s", "token is empty")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token missing"})
			c.Abort()
			return
		}

		claims := &Claims{}
		parsedToken, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !parsedToken.Valid {
			println("token validation failed")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Token missing"})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		c.Next()
	}
}

func getListOfTaskHandler(r *gin.Context) {

	ctx, _ := context.WithTimeout(r.Request.Context(), 30*time.Second)
	listOfTask := []string{}
	for i := range 5 {
		listOfTask = append(listOfTask, fmt.Sprintf("task-%d", i))
	}

	var wg sync.WaitGroup
	valueChannel := make(chan string)
	errChannel := make(chan error)
	var result []string
	var mu sync.Mutex

	for index, task := range listOfTask {
		// do operation concurrently
		wg.Add(1)
		go func(index int, task string) {
			println("processing started")
			defer wg.Done()
			select {
			case <-ctx.Done():
				errChannel <- ctx.Err()
			case <-time.After(time.Duration(time.Duration(index) * time.Second)):
				println("processing.....")
				taskCompletion := fmt.Sprintf("%s - %d -  done", task, index)
				mu.Lock()
				result = append(result, taskCompletion)
				mu.Unlock()
				valueChannel <- taskCompletion

			}
		}(index, task)
	}

	go func() {
		wg.Wait()
		close(valueChannel)
		close(errChannel)
	}()

	for {
		select {
		case err, ok := <-errChannel:
			if ok {
				log.Fatal("error while computing task: %s", err)
			}
		case val, ok := <-valueChannel:
			if ok {
				log.Println("info: %s", val)
			}
		case <-ctx.Done():
			r.JSON(http.StatusRequestTimeout, gin.H{"status": "Timeout"})
			return
		default:
			if len(result) == len(listOfTask) {
				r.JSON(http.StatusOK, gin.H{"data": result})
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {
	fmt.Println("Sample REST Backend")

	gr := gin.Default()

	grGroup := gr.Group("/api/v1/dummy")
	grGroup.Use(requestLogger())

	grGroup.POST("/login", dummyLoginHandler)

	grDashboardGroup := grGroup.Group("/dashboard")
	// grDashboardGroup.Use(tokenValidatorMiddleware())
	grDashboardGroup.GET("/tasks", getListOfTaskHandler)

	gr.Run(":9090")
}
