package main

import (
	"crypto/des"
	"crypto/sha1"
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Exported function without documentation (golint: exported function should have comment)
func ProcessUserData(input string) string {
	return strings.ToLower(input)
}

// Exported type without documentation
type UserRecord struct {
	Name    string
	Email   string
	Age     int
	isAdmin bool // unexported field in exported struct
}

// Exported constant without documentation
const MaxRetries = 5

// error variable not following conventions (should be ErrNotFound)
var NotFoundError = errors.New("not found")

// Unused parameter
func unusedParam(a int, b int, c int) int {
	return a + b
}

// Deeply nested function (cognitive complexity)
func deeplyNested(x int, y int, z int) string {
	if x > 0 {
		if y > 0 {
			if z > 0 {
				if x > y {
					if y > z {
						if x+y > z {
							return "deep"
						} else {
							return "not so deep"
						}
					} else {
						return "medium"
					}
				} else {
					return "shallow"
				}
			} else {
				return "negative z"
			}
		} else {
			return "negative y"
		}
	} else {
		return "negative x"
	}
}

// SQL injection vulnerability
func getUserByName(db *sql.DB, name string) (*sql.Rows, error) {
	query := "SELECT * FROM users WHERE name = '" + name + "'"
	return db.Query(query)
}

// Weak crypto: DES
func encryptWithDES(key []byte) {
	block, err := des.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(block)
}

// Weak crypto: SHA1
func hashWithSHA1(data string) {
	h := sha1.New()
	h.Write([]byte(data))
	fmt.Printf("%x\n", h.Sum(nil))
}

// Using math/rand instead of crypto/rand for security-sensitive operation
func generateToken() string {
	rand.Seed(time.Now().UnixNano())
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, 32)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// HTTP handler with no timeout
func startServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", r.URL.Query().Get("name"))
	})
	http.ListenAndServe(":8080", nil)
}

// Inefficient string concatenation in loop
func buildString(items []string) string {
	result := ""
	for _, item := range items {
		result = result + item + ", "
	}
	return result
}

// Regex compiled inside loop
func matchAll(patterns []string, input string) []bool {
	results := make([]bool, len(patterns))
	for i, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		results[i] = re.MatchString(input)
	}
	return results
}

// Boolean parameter (flag argument anti-pattern)
func fetchData(url string, useCache bool) string {
	if useCache {
		return "cached"
	}
	return url
}

// Naked return with named results
func divide(a, b float64) (result float64, err error) {
	if b == 0 {
		err = fmt.Errorf("division by zero")
		return
	}
	result = a / b
	return
}

// Type assertion without ok check
func unsafeTypeAssert(i interface{}) string {
	return i.(string)
}

// Empty branch
func emptyBranch(x int) {
	if x > 10 {
		// TODO: implement this
	} else {
		fmt.Println("small")
	}
}

// Goroutine leak - no way to stop
func leakyGoroutine() {
	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println("still running...")
		}
	}()
}

// File not closed (resource leak)
func readFileLeaky(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	// Missing defer f.Close()
	buf := make([]byte, 1024)
	n, err := f.Read(buf)
	return buf[:n], err
}

// Error not checked
func writeToFile(path string, data string) {
	f, _ := os.Create(path)
	f.WriteString(data)
	f.Close()
}

// Unnecessary type conversion
func unnecessaryConversion(x int) int {
	return int(x)
}

// Magic numbers
func calculatePrice(quantity int) float64 {
	return float64(quantity) * 19.99 * 1.0825
}

// Redundant nil check
func redundantNilCheck(s *string) string {
	if s != nil {
		if s != nil {
			return *s
		}
	}
	return ""
}

// Function too many parameters
func tooManyParams(a, b, c, d, e, f, g, h string) string {
	return a + b + c + d + e + f + g + h
}

// Inefficient append
func inefficientSlice() []int {
	var s []int
	for i := 0; i < 10000; i++ {
		s = append(s, i)
	}
	return s
}

// Sprintf used just for int to string conversion
func intToString(n int) string {
	return fmt.Sprintf("%d", n)
}

// Should use strconv instead
func badConversion(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

// time.Sleep in a test-like function
func waitForReady() bool {
	for i := 0; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)
		if i == 5 {
			return true
		}
	}
	return false
}

// Unreachable code
func unreachableCode(x int) int {
	return x * 2
	fmt.Println("this will never execute")
	return x * 3
}

// Self-assignment
func selfAssign(x int) int {
	x = x
	return x
}
