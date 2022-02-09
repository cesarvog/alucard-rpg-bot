package main

import (
	"fmt"
	"os"
)

var (
	Token string
	config *configStruct
)

type configStruct struct {
	Token	string 
}

func ReadConfig() error {
	Token = os.Getenv("TOKEN")
	fmt.Printf("token=%s\n", Token)
	return nil
}
