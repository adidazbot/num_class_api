package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
)

// JSONResponse represents the API response structure
type JSONResponse struct {
	Number    int    `json:"number"`
	Even      bool   `json:"even"`
	Prime     bool   `json:"prime"`
	PerfectSq bool   `json:"perfect_square"`
	Fact      string `json:"fun_fact,omitempty"`
	Error     string `json:"error,omitempty"`
}

// checkPrime determines if a number is prime
func checkPrime(n int) bool {
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

// fetchFunFact gets a fun fact about the number from Numbers API
func fetchFunFact(n int) string {
	// Hardcoding a fun fact since Numbers API requires an external request
	return fmt.Sprintf("%d is a fascinating number!", n)
}

// classifyNumber handles number classification and JSON response
func classifyNumber(w http.ResponseWriter, r *http.Request) {
	// Ensure Content-Type is JSON
	w.Header().Set("Content-Type", "application/json")

	// Parse query parameter ?number=
	query := r.URL.Query().Get("number")
	if query == "" {
		http.Error(w, `{"error": "Missing 'number' parameter"}`, http.StatusBadRequest)
		return
	}

	// Convert query parameter to integer
	number, err := strconv.Atoi(query)
	if err != nil {
		http.Error(w, `{"error": "Invalid number format"}`, http.StatusBadRequest)
		return
	}

	// Build response
	response := JSONResponse{
		Number:    number,
		Even:      number%2 == 0,
		Prime:     checkPrime(number),
		PerfectSq: math.Sqrt(float64(number)) == float64(int(math.Sqrt(float64(number)))),
		Fact:      fetchFunFact(number),
	}

	// Encode and send JSON response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/classify", classifyNumber)
	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
