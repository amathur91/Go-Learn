package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	 http.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request){
	 	log.Println("Hello World")
	 	data, err := ioutil.ReadAll(request.Body)
	 	if err != nil {
	 		http.Error(responseWriter, "Bad Request", http.StatusBadRequest)
	 		return
		}
	 	log.Printf("%s", data)
	 	response := "Hello " + string(data) + "\n"
	 	fmt.Fprintf(responseWriter, response)
	 })
	 http.ListenAndServe(":9090", nil)
}
