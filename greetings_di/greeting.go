package greetings_di

import (
	"bytes"
	"fmt"
	"os"
)

func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Arjesh")
}
