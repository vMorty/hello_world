package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Printf("You Suck! give me your first name")
		return
	}
	if isNumeric(os.Args[1]) {
		log.Printf("I'm sorry, but it cannot be a number")
		return
	}

}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func sum(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}
