package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	spanish = "Spanish"
	french  = "French"

	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	englishHelloPrefix = "Hello, "

	suffix = "!"

	finalWord      = "Go!"
	countdownStart = 3
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name + suffix
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}

func main() {
	//fmt.Println(Hello("Arjesh", "Spanish"))
	//log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
	Countdown(os.Stdout)
}

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func Countdown(out io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprintf(out, finalWord)
}
