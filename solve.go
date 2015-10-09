package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/codegangsta/cli"
)

var dict []string // kept as a global to reduce size of recursive call on stack

func solve(base, acc CharList, index int, out chan []string) {
	if !base.withinBounds(acc) { return }
	if index >= len(dict) { return }
	if base.equals(acc) { 
		out <- acc.Components // found a match
		return
	}

	// these are thread safe and should use threads later for speed
	solve(base, acc.addString(dict[index]), index, out) // doesn't change index so that it can use a word multiple times
	solve(base, acc, index + 1, out)
} 

func solveAnagram(c *cli.Context) {
	file, err := ioutil.ReadFile(c.Args().First())

	if err != nil {
		fmt.Errorf("Could not read file |%v|", err)
	}

	dict = strings.Split(string(file), "\n") 
	dict = dict[:len(dict)-1] // last element is empty

	anagramComponents := make(chan []string)

	acc := NewCharList()
	base := acc.addString(baseString)

	go solve(base, acc, 0, anagramComponents)

	for{
		select {
		case strs := <- anagramComponents:
			fmt.Printf("%v\n", strs)
		}
	}
}
