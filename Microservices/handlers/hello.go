package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}
func (h *Hello) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request){
	h.l.Println("Hello World")
	data, err := ioutil.ReadAll(request.Body)
	if  err != nil {
		http.Error(responseWriter, "Bad Request", http.StatusBadRequest)
		return
	}
	log.Printf("%s", data)
	response := "Hello " + string(data) + "\n"
	fmt.Fprintf(responseWriter, response)
}