package read

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadFileAsChars(inputFile string) []string {
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

		for i := 0; i < len(tempString); i++ {
			output = append(output, string(tempString[i]))
		}
		output = append(output, tempString)
	}

	file.Close()
	return output
}

//ReadFileAsInt - Read input file and return []string: one entry for one row
func ReadFileAsStrings(inputFile string) []string {

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
func ReadFileAsInt(sinputFile string) []int {

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
