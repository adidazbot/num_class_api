package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// classifyNumber handles the number classification and returns JSON response.
func classifyNumber(c *gin.Context) {
	numberStr := c.Query("number") // Get number from query params
	number, err := strconv.Atoi(numberStr)

	// Handle invalid inputs
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"number": numberStr,
			"error":  true,
		})
		return
	}

	// Determine number properties
	properties := []string{}
	if isArmstrong(number) {
		properties = append(properties, "armstrong")
	}
	if number%2 == 0 {
		properties = append(properties, "even")
	} else {
		properties = append(properties, "odd")
	}

	// Prepare JSON response
	response := gin.H{
		"number":    number,
		"is_prime":  isPrime(number),
		"is_perfect": isPerfect(number),
		"properties": properties,
		"digit_sum":  digitSum(number),
		"fun_fact":   getFunFact(number),
	}

	// Return a successful response
	c.JSON(http.StatusOK, response)
}

// isPrime checks if a number is prime.
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// isPerfect checks if a number is a perfect number.
func isPerfect(n int) bool {
	sum := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum == n
}

// isArmstrong checks if a number is an Armstrong number.
func isArmstrong(n int) bool {
	sum := 0
	temp := n
	numDigits := len(strconv.Itoa(n))

	for temp > 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(numDigits)))
		temp /= 10
	}

	return sum == n
}

// digitSum calculates the sum of digits of a number.
func digitSum(n int) int {
	sum := 0
	for n != 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

// getFunFact fetches a fun fact about the number from Numbers API.
func getFunFact(n int) string {
	return fmt.Sprintf("%d is a cool number with unique properties, it might just be number of your best day!", n) // Placeholder response
}

func main() {
	// Initialize Gin router
	r := gin.Default()

	// Enable CORS (Allow requests from anywhere)
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	})

	// Define API endpoint
	r.GET("/api/classify-number", classifyNumber)

	// Get the PORT from environment variables (Render assigns a dynamic port)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port for local testing
	}

	// Start the API server
	log.Printf("Server running on port %s...", port)
	err := r.Run(":" + port)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
