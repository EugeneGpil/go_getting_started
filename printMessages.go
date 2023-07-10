package main

import "fmt"

func printMessages(messages map[string]string) {
	for _, message := range messages {
		fmt.Println(message)
	}
}
