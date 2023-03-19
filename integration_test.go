package main

import (
	"fmt"
	"bytes"
	"testing"
	"net/http"
	"io"
	"encoding/json"
	"github.com/henrikkorsgaard/go-file-api/api"
)

/*
	Some of these tests may appear silly, but I'm using an test-driven approach
	that starts with integration-like test (outside in).

	When possible, I will refactor into isolated test cases
*/

type Todo struct {
	Todo string `json:"todo"`
	Due string `json:"due"`
}


func TestAPIServer(t *testing.T){
	port := 3000
	addr := fmt.Sprintf(":%d", port)

	go  http.ListenAndServe(addr, api.API{})// how the fuck do I stop this again

	url := fmt.Sprintf("http://localhost:%d/test-todo", port)

	t.Run("POST /", func(t *testing.T){
		rootUrl := fmt.Sprintf("http://localhost:%d", port)
		req, err := http.NewRequest("POST", rootUrl, bytes.NewBuffer([]byte("")))

		// This should not happen 
		if err != nil {
			t.Fatal("Programmer error!")
		}

		req.Header.Set("Content-Type", "application/json;charset=UTF-8")

		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			t.Logf("Server responded with an error: %v", err)
			t.Fail()
		}
		defer resp.Body.Close()
	
		// We need to make sure that the resource is created!
		if resp.StatusCode != 421 {
			t.Logf("Server responded with status code: %v, expected 421", resp.StatusCode)
			t.Fail()
		}
	})

	t.Run("POST json" , func(t *testing.T){

		marshal_todo, err := json.Marshal(Todo{Todo:"Clean kitchen", Due:"today"})
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(marshal_todo))
		
		// This should not happen 
		if err != nil {
			t.Fatal("Programmer error!")
		}

		req.Header.Set("Content-Type", "application/json;charset=UTF-8")

		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			t.Logf("Server responded with an error: %v", err)
			t.Fail()
		}
		defer resp.Body.Close()
	
		// We need to make sure that the resource is created!
		if resp.StatusCode != 201 {
			t.Logf("Server responded with status code: %v, expected 201", resp.StatusCode)
			t.Fail()
		}
	})

	t.Run("GET json", func(t *testing.T){
		resp, err := http.Get(url)
		defer resp.Body.Close()
		if err != nil {
			t.Logf("Server responded with error: %v ", err)
			t.Fail()
		}
	
		if resp.StatusCode != 200 {
			t.Logf("Server responded with status code: %v", resp.StatusCode)
			t.Fail()
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Logf("Unable to read response body from server: %v ", err)
			t.Fail()
		}

		var todos []Todo
		err = json.Unmarshal(body, &todos)
		if err != nil {
			t.Logf("Unable to unmarshal todos from server: %v ", err)
			t.Fail()
		}
	})
}
