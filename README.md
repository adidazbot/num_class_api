  
# Number Classification API
---

## 📌 The Task: What I Needed to Build  

The goal was to build an API that:  
✅ Accepts a **GET request** with a number parameter  
✅ Returns **mathematical properties** (prime, Armstrong, odd/even, etc.)  
✅ Fetches a **fun fact** using the [Numbers API](http://numbersapi.com/)  
✅ Returns **JSON responses**  
✅ Handles **errors & invalid inputs**  
✅ Is **publicly accessible** and **deployed online**  

---

## 🛠 Step 1: Setting Up the Go Project  

### **1️⃣ Install Go (If You Haven't Already)**  
Before writing any code, I made sure Go was installed:  
```sh
go version
```
If you don’t have Go, install it from [golang.org](https://go.dev/dl/).  

### **2️⃣ Create a New Go Project**  
```sh
mkdir num-class-api && cd num-class-api
go mod init github.com/YOUR_USERNAME/num-class-api
```
This sets up a **Go module** to manage dependencies.  

### **3️⃣ Install Dependencies**  
```sh
go get github.com/gin-gonic/gin
```
I used **Gin**, a lightweight web framework for Go.  

---

## **🖥 Step 2: Writing the API (main.go)**  

Here’s the breakdown of **how the API works**:  

```go
package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Function to check if a number is prime
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

// Function to check if a number is an Armstrong number
func isArmstrong(n int) bool {
	sum, temp, digits := 0, n, len(strconv.Itoa(n))
	for temp > 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(digits)))
		temp /= 10
	}
	return sum == n
}

// Function to fetch fun fact from Numbers API
func getFunFact(n int) string {
	resp, err := http.Get(fmt.Sprintf("http://numbersapi.com/%d?json", n))
	if err != nil {
		return "No fun fact available."
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&data)
	return data["text"].(string)
}

func main() {
	r := gin.Default()

	// Enable CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	})

	r.GET("/api/classify-number", func(c *gin.Context) {
		numStr := c.Query("number")
		num, err := strconv.Atoi(numStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"number": numStr, "error": true})
			return
		}

		properties := []string{}
		if num%2 == 0 {
			properties = append(properties, "even")
		} else {
			properties = append(properties, "odd")
		}
		if isArmstrong(num) {
			properties = append(properties, "armstrong")
		}

		response := gin.H{
			"number":    num,
			"is_prime":  isPrime(num),
			"properties": properties,
			"fun_fact":  getFunFact(num),
		}

		c.JSON(http.StatusOK, response)
	})

	r.Run(":8080")
}
```

✅ **Core Features Implemented:**  
- **Checks if a number is prime**  
- **Detects Armstrong numbers**  
- **Classifies even/odd numbers**  
- **Fetches a fun fact from the Numbers API**  
- **Handles invalid input gracefully**  
- **CORS enabled for public access**  

---

## **🚀 Step 3: Running & Testing the API**  

### **Run the API Locally**  
```sh
go run main.go
```
Open your browser and test:  
```
http://localhost:8080/api/classify-number?number=371
```
Expected Response:  
```json
{
    "number": 371,
    "is_prime": false,
    "properties": ["armstrong", "odd"],
    "fun_fact": "371 is an Armstrong number!"
}
```

---

## **🚢 Step 4: Deploying the API (Railway.app)**  

I chose **Railway.app** for deployment because it’s **fast, free, and supports Go out-of-the-box**.  

### **1️⃣ Push Code to GitHub**  
```sh
git init
git add .
git commit -m "Initial commit"
git branch -M main
git remote add origin https://github.com/adidazbot/num-class-api.git
git push -u origin main
```

### **2️⃣ Deploy on Railway**  
1. Install Railway CLI  
   ```sh
   npm install -g @railway/cli
   ```
2. Login  
   ```sh
   railway login
   ```
3. Create a new project  
   ```sh
   railway init
   ```
4. Deploy 🚀  
   ```sh
   railway up
   ```

🎉 **Success!** My API is now publicly available at:  
```
https://num-class-api.up.railway.app/api/classify-number?number=371
```

---

## **🔥 Challenges & Errors Faced**  

### **1. GitHub Authentication Error**
💥 **Error:** GitHub removed password authentication.  
✅ **Fix:** Used **Personal Access Token (PAT)** instead of a password.

---

### **2. Invalid Input Crashes API**
💥 **Error:** Non-numeric input crashed the API.  
✅ **Fix:** Added proper input validation with `strconv.Atoi()`.

---

### **3. CORS Blocking API Calls**
💥 **Error:** Browser blocked API requests due to CORS.  
✅ **Fix:** Added a **CORS middleware** in Gin.

---

## **🔗 Live API & GitHub Repo**  
🚀 **Live API:** [https://num-class-api.up.railway.app](https://num-class-api.up.railway.app)  
📌 **GitHub Repo:** [https://github.com/adidazbot/num-class-api](https://github.com/adidazbot/num-class-api)  

---

Thank you for visiting this repo! Hope this guide helps.
