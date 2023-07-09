package main

import (
    "fmt"
    "log"

	"github.com/EugeneGpil/greetings"
)

func main() {
    log.SetPrefix("greetings: ")
    // log.SetFlags(0)

    // message, err := greetings.Hello("")
    message, err := greetings.Hello("PES")

    if (err != nil) {
        log.Fatal(err)
    }

    fmt.Println(message)
}
