package main

import (
    // "fmt"
    "log"

	"github.com/EugeneGpil/greetings"
)

func main() {
    log.SetPrefix("greetings: ")
    // log.SetFlags(0)

    names := []string{
        "PES1",
        "PES2",
        "PES3",
    }

    // message, err := greetings.Hello("")
    // message, err := greetings.Hello("PES")
    messages, err := greetings.Hellos(names)

    if (err != nil) {
        log.Fatal(err)
    }

    // fmt.Println(message)
    
    printMessages(messages)
}
