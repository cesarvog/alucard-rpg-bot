package main

import (
	"fmt"
	"net/http"
	"os"
)	
	
func main() {
	err := ReadConfig()
	
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	
	Start()
	go fmt.Println(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
	<-make(chan struct{})
	return
}	
