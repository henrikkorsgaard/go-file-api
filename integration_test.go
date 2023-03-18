package main

import (
	"fmt"
	"testing"
	"net/http"
	"github.com/henrikkorsgaard/go-file-api/api"
)

/*
	Some of these tests may appear silly, but I'm using an test-driven approach
	that starts with integration-like test (outside in).

	When possible, I will refactor into isolated test cases
*/


func TestAPIServerCRUD(t *testing.T){
	port := 3000
	go api.StartAPIServerOnPort(port) // how the fuck do I stop this again

	url := fmt.Sprintf("http://localhost:%d/", port)

	t.Run("GET", func(t *testing.T){
		res, err := http.Get(url)
		defer res.Body.Close()
		if err != nil {
			t.Fatalf("Server responded with error: %v ", err)
		}
	
		if res.StatusCode != 200 {
			t.Fatalf("Server responded with status code: %v", res.StatusCode)
		}
	})

	t.Run("POST" , func(t *testing.T){
		res, err := http.Post(url, "application/json",nil)

		if err != nil {
			t.Fatalf("Server responded with error: %v ", err)
		}
	
		if res.StatusCode != 200 {
			t.Fatalf("Server responded with status code: %v", res.StatusCode)
		}
	})
	
	// TEAR DOWN SHOULD BE STOPPING SERVER!
}


//First test starting a webserver on port?

//second test post JSON

//third test get JSON that matches the posted json

//authentication