package api

import (
	"fmt"
	"strings"
	"io"
	"os"
	"log"
	"net/http"
	"net/url"
	"encoding/json"
	"errors"
)


type API struct {
	//need a file folder config to store files in specific folder
}


func (api API) ServeHTTP(w http.ResponseWriter, r *http.Request){

	if r.Method == "GET" {
		entity := getEntityFromPath(r.URL)
		json, err := readFromFile(entity)
		if err != nil {
			w.WriteHeader(503)
			fmt.Fprintf(w,"503: Service unavailable")
			return
		}

		json = fmt.Sprintf("[%s]",json)
		fmt.Fprintf(w,json)
	} else if r.Method == "POST" {
		
		//we try to get the first entity. That's all we are interested in at this point
		entity := getEntityFromPath(r.URL)
		
		// Return 421: Not Found because the API does not accept post on /
		if entity == "" {
			w.WriteHeader(421)
			fmt.Fprintf(w,"421: Cannot accept POST to /")
			return
		}
		
		//Maybe just try to demarshal and catch errors
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		var jsonObj interface{}
		err = json.Unmarshal(body, &jsonObj)
		if err != nil {
			w.WriteHeader(507)
			fmt.Fprintf(w,"507: Unable to store entity due to unknown error")
			return
		}

		//At this point we trust that the incoming string is parsable by the json library, hence we can store it!
		err = writeToFile(entity, string(body))
		
		if err != nil {
			w.WriteHeader(507)
			fmt.Fprintf(w,"507: Unable to store entity due to unknown error")
			return
		} else {
			w.WriteHeader(201)
			fmt.Fprintf(w,"201: Entity created")
			return
		}

	} else {
		w.WriteHeader(400)
		fmt.Fprintf(w,"400")
	}
}

func getEntityFromPath(url *url.URL) (entity string) {

	splitFunc := func(c rune) bool {
		return c == '/'
	}

	elements := strings.FieldsFunc(url.Path, splitFunc)

	if len(elements) > 0 {
		entity = elements[0]
	}
	return 
}

func writeToFile(entity string, json string) error {
	filename := fmt.Sprintf("%s.henrik", entity)

	fileNotExisting := false;

	if _, unerr := os.Stat(filename); errors.Is(unerr, os.ErrNotExist) {
		fileNotExisting = true
	}

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	} 
	defer file.Close()

	line := fmt.Sprintf(",\n%s", json) //because we need to add , to fix the array
	if fileNotExisting {
		line = fmt.Sprintf("%s", json) //because we need to add , to fix the array
	}
	_, err = file.WriteString(line)
	return err
}

func readFromFile(entity string) (string, error){
	filename := fmt.Sprintf("%s.henrik", entity)
	data, err := os.ReadFile(filename)
 
	return string(data), err
} 