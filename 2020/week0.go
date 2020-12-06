// That time of year again!!! Advent of Code 2020 here we go!
package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

// fetches input from aoc servers
func fetchInput(day int) (filename string, err error) {
	if day > 25 {
		return "", errors.New("there are only 25 days of puzzles!")
	}

	// need to add cookie or it wont give input
	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/2020/day/%d/input", day), nil)
	if err != nil {
		return "", err
	}
	// TODO: see prev comment

	resp, err := http.Get(fmt.Sprintf("https://adventofcode.com/2020/day/%d/input", day))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	filename = fmt.Sprintf("day%dinput.txt", day)
	err = ioutil.WriteFile(filename, data, os.ModePerm|os.ModeAppend)
	if err != nil {
		return "", err
	}

	return filename, nil
}

// day 1
func expenseReport() {
	f, err := fetchInput(1)
	if err != nil {
		panic(err)
	}
	input, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	b := bufio.NewScanner(input)
	// we dont care about the error, this is just an optimization, append
	// will allocate a new slice every iteration , but its neither an error,
	// or anything to panic over.
	tb, _ := ioutil.ReadFile(f)
	cnt := bytes.Count(tb, []byte("\n"))

	var nums = make([]int64, cnt)
	for b.Scan() && b.Err() == nil {
		n, err := strconv.ParseInt(b.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		nums = append(nums, n)
	}
	n0, n1, err := unsum(2020, nums)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d\t%d\n", n0, n1)

}

// called unsum because I couldn't think of the name for the "factors" of
// addition.
// unsum iterates through pool and returns the first two numbers whos sum is
// equal to val, or an error if none are found.
func unsum(val int64, pool []int64) (int64, int64, error) {
	for i := range pool[len(pool)/2:] {
		for j := range pool[:len(pool)/2] {
			if pool[i]+pool[j] == val {
				return pool[i], pool[j], nil
			}
		}

	}
	return 0, 0, fmt.Errorf("could not find any candidates whos sum is %d", val)

}

func main() {
	expenseReport()
}
