package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/codegangsta/cli"
)

func solve(base, acc CharList, dict []string, index int, out chan []string) {
	if !base.withinBounds(acc) { return }
	if index >= len(dict) { return }
	if base.equals(acc) { 
		out <- acc.Components
		return
	}

	solve(base, acc.addString(dict[index]), dict, index, out) // doesn't change index so that it can use a word multiple times
	solve(base, acc, dict, index + 1, out)
} 

func solveAnagram(c *cli.Context) {
	file, err := ioutil.ReadFile(c.Args().First())

	if err != nil {
		fmt.Errorf("Could not read file |%v|", err)
	}

	dict := strings.Split(string(file), "\n") // last element is empty

	channel := make(chan []string)

	acc := NewCharList()
	base := acc.addString(baseString)

	go solve(base, acc, dict[:len(dict)-1], 0, channel)

	for{
		select {
		case strs := <- channel:
			fmt.Printf("we got one %v\n", strs)
			/*for str := range strs {
				fmt.Printf("%s ", str)
			}
			fmt.Println()*/
		}
	}
}