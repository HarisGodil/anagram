package main 

import (
	"bufio"
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func fileToList(c *cli.Context) {

	file, err := os.Open(c.Args().First())
	if err != nil {
		fmt.Errorf("|%v|", err)
	}
	defer file.Close()

	acc := NewCharList()
	base := acc.addString(baseString)

	scanner := bufio.NewScanner(file)

	prevString := "" // to prevent duplicates in dictionary

	for scanner.Scan() {
		line := scanner.Text()

		if line == prevString {
			continue
		}

		entry := acc.addString(line)

		if base.withinBounds(entry) {

			prevString = line

			fmt.Printf("%s\n", entry.Components[0])
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Errorf("|%v|", err)
	}
}