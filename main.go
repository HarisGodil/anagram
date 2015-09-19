package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const md5hash = "4624d200580677270a54ccff86b9610e"
const baseString = "poultry outwits ants"

func main() {
	fileToList("/home/haris/RubyAnagram/output.txt")
}

func fileToList(filePath string) []charList {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	acc := NewCharList()
	base := acc.addString(baseString)

	scanner := bufio.NewScanner(file)

	dictionary := []charList{}

	prevString := "" // to prevent duplicates

	for scanner.Scan() {
		line := scanner.Text()

		if line == prevString {
			continue
		}

		entry := acc.addString(line)

		if base.withinBounds(entry) {

			prevString = line

			dictionary = append(dictionary, entry)
			fmt.Printf("%s\n", entry.components[0])
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return dictionary
}
