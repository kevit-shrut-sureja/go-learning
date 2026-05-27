package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
	Name     string `validate:"required"`
}

var validate *validator.Validate

func main() {
	validate = validator.New()

	// Create a pool of requests (simulating thousands of requests)
	numRequests := 10000000
	wg := sync.WaitGroup{}
	wg.Add(numRequests)

	start := time.Now()

	// Simulate validating many requests concurrently
	for i := 0; i < numRequests; i++ {
		go func() {
			defer wg.Done()

			user := User{
				Email:    "test@example.com",
				Password: "password123",
				Name:     "Test User",
			}

			err := validate.Struct(user)
			if err != nil {
				fmt.Println("Validation error:", err)
			}
		}()
	}

	wg.Wait()

	duration := time.Since(start)
	fmt.Printf("Validated %d requests in %v\n", numRequests, duration)
}
