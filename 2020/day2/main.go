package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	raw := readInoutFile()
	pwd := make([]passwordAndRules, len(raw))
	for i, s := range raw{
		pwd[i] = parse(s)
	}
	count := 0
	for _, p := range pwd{
		if part2Isvalid(p) {
			count++
		}
	}
	fmt.Println(count)

}
func part1Isvalid(pwdr passwordAndRules)  bool{
	count := strings.Count(pwdr.password, pwdr.character)
	if count >= pwdr.min && count <= pwdr.max{
		return true
	}
	return false
}
func part2Isvalid(pwdr passwordAndRules)  bool{
	charAtMin := string(pwdr.password[pwdr.min - 1])
	charAtMax := string(pwdr.password[pwdr.max - 1])

	if charAtMin == pwdr.character && charAtMax != pwdr.character {
		return true
	}
	if charAtMin != pwdr.character && charAtMax == pwdr.character {
		return true
	}
	return false
}

type passwordAndRules struct{
	min int
	max int
	character string
	password string
}

// Input format
// "5-6 d: ddddddb"
func parse(data string) passwordAndRules{

	spaceSplit := strings.Split(data, " ")
	minMax := strings.Split(spaceSplit[0], "-")
	char := spaceSplit[1][0]
	pass := spaceSplit[2]
	min, _ := strconv.Atoi(minMax[0])
	max, _ := strconv.Atoi(minMax[1])
	return passwordAndRules{
		min: min,
		max: max,
		character: string(char),
		password: pass,
	}
}

func readInoutFile() []string {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	return lines
}