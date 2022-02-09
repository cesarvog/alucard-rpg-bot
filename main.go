package main

import (
	"fmt"
)	
	
func main() {
	err := ReadConfig()
	
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	
	Start()
	
	<-make(chan struct{})
	return
}	


