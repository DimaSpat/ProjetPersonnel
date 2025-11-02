package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Print("Please username: ")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input", err)
		return
	}

	input = strings.TrimSuffix(input, "\n")
	currentTime := time.Now()
	dateString := fmt.Sprintf("%d / %s / %d", currentTime.Year(), currentTime.Month(), currentTime.Day())
	
	fmt.Println("Hello " + input + ", current date is " + dateString)
}
