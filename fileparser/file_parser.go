package fileparser

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	// DataDelimiter is the character(s) used to separate each field in a data line
	DataDelimiter = "::"
)

// RowAction is a function that accepts a slice of field values for each row and allow capture of fields
type RowAction func(fields []string) error

// ParseFile accepts a file path to read from and expected field count, it loops through the file line by line.
// For each line it validates the number of fields matches the field count and runs the action function.
func ParseFile(filePath string, fieldCount int, action RowAction) error {
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Fails to open file %s: %s\n", filePath, err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for {
		if ok := scanner.Scan(); !ok {
			if err := scanner.Err(); err != nil {
				log.Printf("Fails to scan file: %s\n", err)
				return err
			}
			// exit the loop when either error or end of file
			break
		}

		line := scanner.Text()
		fields := strings.Split(line, DataDelimiter)
		if len(fields) != fieldCount {
			msg := fmt.Sprintf("Missing fields on line: %s", line)
			log.Println(msg)
			return errors.New(msg)
		}

		if err := action(fields); err != nil {
			log.Println("action fails with error, exiting...")
			return err
		}
	}

	return nil
}
