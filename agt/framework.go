package agt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"ia04/comsoc"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// Respond as JSON
func respondJSON(w http.ResponseWriter, statuscode int, value any) {
	w.WriteHeader(statuscode)
	w.Header().Set("Content-Type", "application/json")
	serial, _ := json.Marshal(value)
	w.Write(serial)
	log.Printf("[JSON] %s", string(serial[:]))
}

// Function to use to send JSON
type Response = func(statuscode int, value any)

// Create a request handler by deserialize the request and send it to the inner function
func route[Request any](method string, do func(Request, Response) error) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check if method is valid
		log.Printf("[%s] %s", r.Method, r.URL)
		if r.Method != method {
			respondJSON(w, http.StatusMethodNotAllowed, ResponseMessage{Message: "MethodNotAllowed"})
			return
		}
		// Deserialize json
		var request Request
		buf := new(bytes.Buffer)
		_, err := buf.ReadFrom(r.Body)
		if err != nil {
			respondJSON(w, http.StatusBadRequest, ResponseMessage{Message: "The request could not be read"})
			return
		}
		err = json.Unmarshal(buf.Bytes(), &request)
		if err != nil {
			respondJSON(w, http.StatusBadRequest, ResponseMessage{Message: "Invalid json body"})
			return
		}
		// Run function
		err = do(request, func(statuscode int, value any) {
			respondJSON(w, statuscode, value)
		})
		if err != nil {
			if httpError, ok := err.(comsoc.HTTPError); ok {
				respondJSON(w, httpError.Code, ResponseMessage{Message: httpError.Message})
			} else {
				respondJSON(w, http.StatusInternalServerError, ResponseMessage{Message: "Internal Server Error"})
			}
		}
	}
}

// Send request to a remote server and parse result as json
func request[R any](url string, req any) (resp R, err error) {
	data, err := json.Marshal(req)
	if err != nil {
		return resp, err
	}
	res, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return resp, err
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(res.Body)
	if err != nil {
		return resp, err
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		var errResp ResponseMessage
		err = json.Unmarshal(buf.Bytes(), &errResp)
		if err != nil {
			return resp, err
		}
		return resp, comsoc.HTTPErrorf(res.StatusCode, errResp.Message)
	} else {
		err = json.Unmarshal(buf.Bytes(), &resp)
	}
	return resp, err
}

func WaitAvailable(url string, duration time.Duration) bool {
	start := time.Now()
	for {
		_, err := http.Post(url, "application/json", bytes.NewBuffer([]byte{}))
		if err == nil {
			return true
		}
		if duration > time.Now().Sub(start) {
			return false
		}
	}
}

// Parse an int or exit program
func ParseInt(arg string, name string) int {
	i, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s (%s) is not an int value\n", arg, name)
		os.Exit(1)
	}
	return i
}

// Parse an string list or exit program
func ParseStringList(arg string, name string) []string {
	return strings.Split(arg, ",")
}

// Parse an Alternative list or exit program
func ParseAlternatives(arg string, name string) []comsoc.Alternative {
	strings := ParseStringList(arg, name)
	ints := make([]comsoc.Alternative, len(strings))
	for i, part := range strings {
		ints[i] = comsoc.Alternative(ParseInt(part, fmt.Sprintf("%s[%d]", name, i)))
	}
	return ints
}

// Parse an int list or exit program
func ParseIntList(arg string, name string) []int {
	strings := ParseStringList(arg, name)
	ints := make([]int, len(strings))
	for i, part := range strings {
		ints[i] = ParseInt(part, fmt.Sprintf("%s[%d]", name, i))
	}
	return ints
}

// Parse time or exit program
func ParseTime(arg string, name string) time.Time {
	timestamp, err := time.Parse(time.RFC3339, arg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse %s: %s\n", name, err)
		os.Exit(1)
	}
	return timestamp
}
