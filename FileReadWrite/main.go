package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	fmt.Print("Enter filename to read/edit (e.g., data.txt): ")
	scanner.Scan()
	filename := scanner.Text()
	filename = strings.TrimSpace(filename)

	if filename == "" {
		fmt.Println("No filename entered. Exiting.")
		return
	}

	fmt.Printf("\n--- Ready to edit '%s' ---\n", filename)

	for {
		fmt.Println("\nOptions:")
		fmt.Println("1. Read file contents")
		fmt.Println("2. Write to file (will overwrite!)")
		fmt.Println("3. Exit")
		fmt.Print("Enter choice: ")

		scanner.Scan()
		input := scanner.Text()

		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid choice, please enter a number.")
			continue
		}

		switch choice {
		case 1:
			fmt.Printf("\n--- Reading file '%s' ---\n", filename)
			contentRead, err := os.ReadFile(filename)
			if err != nil {
				if os.IsNotExist(err) {
					fmt.Println("File doesn't exist yet. Try writing to it first.")
				} else {
					fmt.Printf("Error reading file '%s': %v\n", filename, err)
				}
			} else {
				// File was read successfully
				fmt.Println("File exists. Current content:")
				fmt.Println("----------------------------------")
				fmt.Println(string(contentRead))
				fmt.Println("----------------------------------")
			}
		case 2:
			// --- This is the "Writing" logic from your example ---
			fmt.Print("\nEnter the new content to write: ")
			scanner.Scan() // Use scanner to read the content line
			contentToWrite := scanner.Text()

			fmt.Printf("\n--- Writing to file '%s' ---\n", filename)
			// os.WriteFile creates the file if it doesn't exist, or overwrites it.
			err = os.WriteFile(filename, []byte(contentToWrite), 0644)
			if err != nil {
				fmt.Printf("Error writing to file '%s': %v\n", filename, err)
			} else {
				fmt.Printf("Successfully wrote %d bytes to '%s'.\n", len(contentToWrite), filename)
			}
		case 3:
			fmt.Println("Exiting.")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice, please select 1, 2, or 3.")
		}
	}
}

// NOTE: For larger files, you would use os.Open and bufio.NewScanner
// or bufio.NewReader for memory-efficient, line-by-line reading.
// os.WriteFile and os.ReadFile are best for smaller files (a few megabytes).
