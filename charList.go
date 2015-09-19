package main

type charList struct {
	chars      []int //length 26, chars[0] is # of 'a's, chars[1] is # of 'b's, etc
	components []string
}

func NewCharList() charList {
	return charList{
		chars:      make([]int, 26),
		components: make([]string, 0, 15),
	}
}

// takes in a string, and will return a new charList with that string added
func (c charList) addString(add string) charList {
	newChars := generateInts(add)

	for i := 0; i < 26; i++ {
		newChars[i] += c.chars[i]
	}

	return charList{
		chars:      newChars,
		components: append(c.components, add),
	}
}

func (c charList) withinBounds(other charList) bool {
	for i, val := range c.chars {
		if val < other.chars[i] {
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

		if ascii == 32 || ascii == 39 { // space
			continue
		}

		ascii = ascii - 97
		chars[ascii]++
	}

	return chars
}
