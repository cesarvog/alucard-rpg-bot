package main

import (
	"fmt"
	"net/http"
	"os"
)	

var (
	token = os.Getenv("TOKEN")
	httpPort = os.Getenv("PORT")
)
	
func main() {
	if token == "" {
		fmt.Printf("TOKEN not defined")
		return
	}

	if httpPort == ":" {
		fmt.Printf("PORT not defined")
		return
	}

	Start()
	http.HandleFunc("/", wakeUp)
	go fmt.Println(http.ListenAndServe(":"+os.Getenv("PORT"), nil))

	<-make(chan struct{})
	return
}	

func wakeUp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "I am awaked Milord")
}
