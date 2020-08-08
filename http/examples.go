package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// GetRequest Native http GET request
func GetRequest() {
	url := "http://swapi.dev/api/people/1/"
	resp, err := http.Get(url)
	checkError(err)
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	checkError(err)
	log.Println(jsonPrettyPrint(string(data)))
}

// PostRequest Native http Post request
func PostRequest() {
	data, err := json.Marshal(map[string]string{
		"username": "go user",
		"email":    "go@go.com",
	})
	url := "https://httpbin.org/post"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	checkError(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)
	fmt.Println(string(body))
}

// UsingHTTPManager ...
func UsingHTTPManager() {
	b, err := json.Marshal(map[string]string{
		"username": "go user",
		"email":    "go@go.com",
	})
	checkError(err)
	req := &Request{
		Method:  Post,
		BaseURL: "https://httpbin.org/post",
		Body:    b,
	}
	resp, err := PerformRequest(req)
	checkError(err)
	fmt.Println(resp)
}

// UsingHTTPManagerCustomClient ...
func UsingHTTPManagerCustomClient() {
	c := &http.Client{
		Timeout:   time.Duration(5 * time.Second),
		Transport: http.DefaultTransport,
	}
	cc := Client{
		HTTPClient: c,
	}
	b, err := json.Marshal(map[string]string{
		"username": "go user",
		"email":    "go@go.com",
	})
	checkError(err)
	req := &Request{
		Method:  Post,
		BaseURL: "https://httpbin.org/post",
		Body:    b,
	}
	resp, err := cc.PerformRequest(req)
	checkError(err)
	fmt.Println(resp)
}

// UsingHTTPManagerGet ...
func UsingHTTPManagerGet() {
	req := &Request{
		Method:  Post,
		BaseURL: "http://swapi.dev/api/people/1/",
	}
	resp, err := PerformRequest(req)
	checkError(err)
	log.Println(jsonPrettyPrint(string(resp.Body)))
}

// UsingHTTPManagerPatch ...
func UsingHTTPManagerPatch() {
	b, err := json.Marshal(map[string]string{
		"username": "go user",
		"email":    "go@go.com",
	})
	checkError(err)
	req := &Request{
		Method:  Patch,
		BaseURL: "https://httpbin.org/patch",
		Body:    b,
	}
	resp, err := PerformRequest(req)
	checkError(err)
	fmt.Println(resp)
}

// UsingHTTPManagerPut ...
func UsingHTTPManagerPut() {
	b, err := json.Marshal(map[string]string{
		"username": "go user",
		"email":    "go@go.com",
	})
	checkError(err)
	req := &Request{
		Method:  Put,
		BaseURL: "https://httpbin.org/put",
		Body:    b,
	}
	resp, err := PerformRequest(req)
	checkError(err)
	fmt.Println(resp)
}

// UsingHTTPManagerDelete ...
func UsingHTTPManagerDelete() {
	req := &Request{
		Method:  Delete,
		BaseURL: "https://httpbin.org/delete",
	}
	resp, err := PerformRequest(req)
	checkError(err)
	fmt.Println(resp)
}
