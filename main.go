package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func status(w http.ResponseWriter, r *http.Request) {
	log.Println("Accessed status() function")
	str := `{
        {"result":"OK - healthy"},
        status=200,
        mimetype='application/json'
    }`

	var obj map[string]interface{}
	json.Unmarshal([]byte(str), &obj)

	fmt.Fprint(w, str)

}

func author(w http.ResponseWriter, r *http.Request) {
	log.Println("Accessed author() function")
	str := `{"Author Name":"Gagan Kalra",
    "Github": "@igagankalra(https://github.com/igagankalra)",
    "Twitter": "@igagankalra(https://twitter.com/igagankalra)"
}`

	// var obj map[string]interface{}
	// json.Unmarshal([]byte(str), &obj)

	fmt.Fprint(w, str)

}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	log.Println("Accessed Hello world! function")
	fmt.Fprintf(w, "Hello World, from Gagan!")
}

func metrics(w http.ResponseWriter, r *http.Request) {
	log.Println("Accessed metrics() function")
	str := `{
        {
            "status": "Success",
            "Code": 0,
            "data": {
                "UsrCount":120,
                "UserCountActive":23
            }
        },
        status=200,
        mimetype='application/json'

    }`

	var obj map[string]interface{}
	json.Unmarshal([]byte(str), &obj)

	fmt.Fprint(w, str)
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/metrics", metrics)
	http.HandleFunc("/status", status)
	http.HandleFunc("/author", author)
	http.ListenAndServe(":6111", nil)

}
