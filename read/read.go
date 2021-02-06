package read

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

//ReadFileAsInt - Read input file and return []string
func ReadFileAsString(inputFile string) []string {

	file, err := os.Open(inputFile)
	output := []string{}

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		var tempString = scanner.Text()
		if err == nil {
		}
		output = append(output, tempString)
	}

	file.Close()
	return output
}

//ReadFileAsInt - Read input file and return []int
func ReadFileAsInt(inputFile string) []int {

	file, err := os.Open(inputFile)
	output := []int{}

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		var tempString = scanner.Text()
		if err == nil {
		}
		tempInt, _ := strconv.Atoi(tempString)
		output = append(output, tempInt)
	}

	file.Close()
	return output
}
