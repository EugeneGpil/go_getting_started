package main

import (
    "fmt"

	"github.com/EugeneGpil/greetings"
)

func main() {
    message := greetings.Hello("PES")

    fmt.Println(message)
}
