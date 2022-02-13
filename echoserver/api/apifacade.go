package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	mux "github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var message Message
	json.Unmarshal(reqBody, &message)

	fmt.Println("Endpoint Hit: echo\nMessage: ", string(message.Text))
	json.NewEncoder(w).Encode(message)
}

func handleRequests(port int32) {
	myRouter := mux.NewRouter().StrictSlash(true)

	//POST - Echo endpoint with message body
	myRouter.HandleFunc(echoEndpoint, homePage).Methods(POST)

	var newPort string = fmt.Sprintf("%s%d", ":", port)
	log.Fatal(http.ListenAndServe(newPort, myRouter))
}

func StartServer(port int32) {
	fmt.Println("Running Server in port", port)
	handleRequests(port)
}
