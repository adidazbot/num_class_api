package main

import (
	"encoding/json"  // For JSON encoding and decoding
	"fmt"            // For formatted I/O
	"math"           // For mathematical operations
	"net/http"       // For handling HTTP requests
	"strconv"        // For converting strings to numbers

	"github.com/gorilla/mux" // Router package for handling routes
	"github.com/rs/cors"     // Middleware for handling CORS
)

// Response struct defines the JSON response format for the API
type Response struct {
	Number     int      `json:"number"`    // The input number
	IsPrime    bool     `json:"is_prime"`  // Whether the number is prime
	IsPerfect  bool     `json:"is_perfect"`// Whether the number is a perfect number
	Properties []string `json:"properties"`// Properties (odd/even, armstrong)
	DigitSum   int      `json:"digit_sum"` // Sum of the digits of the number
	FunFact    string   `json:"fun_fact"`  // Fun fact fetched from Numbers API
	Error      bool     `json:"error,omitempty"` // Error field (only appears when there's an error)
}

// Function to check if a number is prime
func isPrime(n int) bool {
	if n < 2 {
		return false // Numbers < 2 are not prime
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false // If divisible by any number, it's not prime
		}
	}
	return true // If no divisors, it's prime
}

// Function to check if a number is a perfect number
// A perfect number is a number whose sum of divisors (excluding itself) equals the number itself
func isPerfect(n int) bool {
	if n < 2 {
		return false
	}
	sum := 1 // Start with 1 as it's a divisor of all numbers
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			sum += i // Add divisors to sum
		}
	}
	return sum == n // If sum of divisors equals the number, it's perfect
}

// Function to check if a number is an Armstrong number
// An Armstrong number (narcissistic number) is a number that is equal to the sum of its own digits each raised to the power of the number of digits
func isArmstrong(n int) bool {
	temp, sum := n, 0
	digits := len(strconv.Itoa(n)) // Count number of digits in the number

	for temp > 0 {
		digit := temp % 10                          // Extract the last digit
		sum += int(math.Pow(float64(digit), float64(digits))) // Add digit^digits to sum
		temp /= 10                                  // Remove last digit
	}

	return sum == n // If sum of powered digits equals original number, it's Armstrong
}

// Function to calculate the sum of digits of a number
func digitSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10 // Extract last digit and add to sum
		n /= 10       // Remove last digit
	}
	return sum // Return total sum of digits
}

// Function to fetch a fun fact from Numbers API
func getFunFact(number int) string {
	url := fmt.Sprintf("http://numbersapi.com/%d/math", number) // Format API URL
	resp, err := http.Get(url) // Make API request
	if err != nil {
		return "Could not fetch fun fact" // Error handling
	}
	defer resp.Body.Close() // Ensure response body is closed after function exits

	var fact string
	_, err = fmt.Fscan(resp.Body, &fact) // Read response into fact variable
	if err != nil {
		return "Could not parse fun fact"
	}

	return fact // Return the fun fact
}

// API Handler to classify a number and return its properties in JSON format
func classifyNumber(w http.ResponseWriter, r *http.Request) {
	// Get number from query parameter
	query := r.URL.Query().Get("number")
	num, err := strconv.Atoi(query) // Convert string to integer

	if err != nil {
		// If conversion fails, return a 400 Bad Request response
		http.Error(w, `{"number":"`+query+`","error":true}`, http.StatusBadRequest)
		return
	}

	// Determine number properties
	properties := []string{}
	if num%2 == 0 {
		properties = append(properties, "even") // Number is even
	} else {
		properties = append(properties, "odd") // Number is odd
	}
	if isArmstrong(num) {
		properties = append(properties, "armstrong") // Add Armstrong if applicable
	}

	// Create response struct
	response := Response{
		Number:     num,
		IsPrime:    isPrime(num),
		IsPerfect:  isPerfect(num),
		Properties: properties,
		DigitSum:   digitSum(num),
		FunFact:    getFunFact(num),
	}

	// Set response header to JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response) // Encode response struct as JSON
}

func main() {
	// Initialize Gorilla Mux Router
	r := mux.NewRouter()

	// Define API route for number classification
	r.HandleFunc("/api/classify-number", classifyNumber).Methods("GET")

	// Enable CORS (allows requests from other origins)
	handler := cors.AllowAll().Handler(r)

	// Start the HTTP server on port 8080
	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", handler)
}

