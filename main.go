package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/cli"
)

const md5hash = "4624d200580677270a54ccff86b9610e"
const baseString = "poultry outwits ants"

///////////////////////////////////////////////////////
// USAGE:
// go build
// ./anagram reduce [wordlist file path] > [reduced wordlist filepath] 
// ./anagram find [reduced wordlist prefered, but any wordlist works] 
// For example
// ./anagram reduce /home/haris/Downloads/wordlist > o.txt
// ./anagram find o.txt
///////////////////////////////////////////////////////

func main() {
	app := cli.NewApp()
	app.Name = "anagram"
	app.Usage = "find anagrams"
	app.Commands = []cli.Command{reduce, find}
	app.Run(os.Args)
}

var reduce = cli.Command{
	Name:    "reduce",
	Usage:   "reduces the dictionary",
	Aliases: []string{"r"},
	Action:  fileToList,
}

var find = cli.Command{
	Name:    "find",
	Usage:   "finds the anagram",
	Aliases: []string{"f"},
	Action:  fileToList,
}

	//fileToList("/home/haris/Downloads/wordlist")

func fileToList(c *cli.Context) {

	file, err := os.Open(c.Args().First())
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	acc := NewCharList()
	base := acc.addString(baseString)

	scanner := bufio.NewScanner(file)

	prevString := "" // to prevent duplicates
	counter := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == prevString {
			continue
		}

		//fmt.Printf("%d\n", counter)
		counter ++
		entry := acc.addString(line)

		if base.withinBounds(entry) {

			prevString = line

			fmt.Printf("%s\n", entry.components[0])
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
