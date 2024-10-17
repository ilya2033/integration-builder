package main

import "fmt"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func log(message string) {
	fmt.Println(message)
}
