package main

import (
	//"fmt"
)

type CharList struct {
	Chars      []int //length 26, chars[0] is # of 'a's, chars[1] is # of 'b's, etc
	Components []string
}

func NewCharList() CharList {
	return CharList{
		Chars:      make([]int, 26),
		Components: make([]string, 0, 15),
	}
}

// takes in a string, and will return a new charList with that string added
func (c CharList) addString(add string) CharList {
	newChars := generateInts(add)

	for i := 0; i < 26; i++ {
		newChars[i] += c.Chars[i]
	}

	return CharList{
		Chars:      newChars,
		Components: append(c.Components, add),
	}
}

func (c CharList) withinBounds(other CharList) bool {

	for i, val := range c.Chars {
		if val < other.Chars[i] {
			return false
		}
	}
	return true
}

func (c CharList) equals(other CharList) bool {

	for i, val := range c.Chars {
		if val != other.Chars[i] {
			return false
		}
	}
	return true
}

// assumes that the string passed in is all lowercase and only has alphabetical (and space) characters
func generateInts(base string) []int {
	chars := make([]int, 26)

	for i, len := 0, len(base); i < len; i++ {
		ascii := int(base[i])

		if ascii == 32 || ascii == 39 { // space or apostrophe ignored
			continue
		}
		//fmt.Printf("%c %d\n", ascii, ascii)

		ascii = ascii - 97
		if ascii < 0 || ascii > 26 {
			chars[0] = 99999; // used to prevent words with special characters from being used
							  // should be changed to something that is more clear
			continue;
		}

		chars[ascii]++
	}

	return chars
}
