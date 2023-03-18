package api

import (
	"fmt"
	"log"
	"net/http"
)

func StartAPIServerOnPort(port int){
	addr := fmt.Sprintf(":%d", port)
	http.HandleFunc("/", apiRouter)
	log.Fatal(http.ListenAndServe(addr, nil))
}

// https://restfulapi.net/http-status-codes/
func apiRouter(w http.ResponseWriter, r *http.Request){
	
	if r.Method == "GET" {
		fmt.Fprintf(w,"")
	} else if r.Method == "POST" {
		fmt.Fprintf(w,"")
	} else {
		w.WriteHeader(400)
	}

	//GET
	//POST -> if t
}