package main

import (
	"log"
)

// bad comment
func f() {

}

func main() {
	log.Printf("invalid format string: %s", 1)
}
