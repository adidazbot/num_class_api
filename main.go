package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// Function to check if a number is prime
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Function to check if a number is perfect (sum of its proper divisors equals the number)
func isPerfect(n int) bool {
	if n < 1 {
		return false
	}
	sum := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum == n
}

// Function to check if a number is an Armstrong number
func isArmstrong(n int) bool {
	sum := 0
	temp := n
	digits := len(strconv.Itoa(n))

	for temp != 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(digits)))
		temp /= 10
	}
	return sum == n
}

// Function to classify number properties
func classifyProperties(n int) []string {
	properties := []string{}

	if isArmstrong(n) {
		properties = append(properties, "armstrong")
	}

	if n%2 == 0 {
		properties = append(properties, "even")
	} else {
		properties = append(properties, "odd")
	}

	return properties
}

// Function to sum the digits of a number
func sumDigits(n int) int {
	n = int(math.Abs(float64(n))) // Ensure positive for digit sum
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

// Function to get a fun fact from the Numbers API
func getFunFact(n int) string {
	return fmt.Sprintf("%d is a cool number with unique properties!", n)
}

// API Handler Function
func classifyNumber(c *gin.Context) {
	numberStr := c.Query("number")

	// Check if number is missing
	if numberStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"number": "missing",
			"error":  true,
		})
		return
	}

	// Validate integer input (Reject floating points)
	if strings.Contains(numberStr, ".") {
		c.JSON(http.StatusBadRequest, gin.H{
			"number": numberStr,
			"error":  true,
		})
		return
	}

	// Convert input to integer
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"number": numberStr,
			"error":  true,
		})
		return
	}

	// Process valid number
	result := map[string]interface{}{
		"number":     number,
		"is_prime":   isPrime(number),
		"is_perfect": isPerfect(number),
		"properties": classifyProperties(number),
		"digit_sum":  sumDigits(number),
		"fun_fact":   getFunFact(number),
	}

	c.JSON(http.StatusOK, result)
}

// Main function to start the API server
func main() {
	r := gin.Default()

	// Enable CORS to allow cross-origin requests
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	})

	// Define the API endpoint
	r.GET("/api/classify-number", classifyNumber)

	// Start server on port 8080
	r.Run(":8080")
}
