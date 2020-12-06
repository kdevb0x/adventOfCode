package main

import (
	"fmt"
	"io/ioutil"
)

var link = "https://adventofcode.com/2017/day/1/input"

/*
After trying for many hours to get the input via http request I have come to the realization
that this was won't work. Using these functions the resp comes back blank. Using curl,
the resp come back along the lines of: " Pleasee login, the input is different for every user."
Looking at the page source reveals only the input, so it must be ajax or something.
Since I am unsure of how to continue, I am just going to skip this part, and hard-code
the input to the file.

func main() {
	var respdata []byte
	err := getDataFromHTTP(link, respdata)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(respdata))

	if err := writeToFile(deriveFileNameFromLink(link), respdata); err != nil {
		log.Println(err)
	}
}
*/

var input, _ = ioutil.ReadFile("2017day1input.txt")

func sumofMatches(input []byte) int {
	var vals []int
	running := 0
	last := 0
	for _, val := range input {
		if int(val) == last {
			vals = append(vals, int(val))
			running += int(val)
		}
		last = int(val)
	}
	if vals[0] == vals[len(vals)-1] {
		vals = append(vals, vals[0])
	}
	var tots int
	for _, cur := range vals {
		tots += cur
	}
	return tots
}

func main() {
	total := sumofMatches(input)
	fmt.Printf("total: %d\n", total)
}
