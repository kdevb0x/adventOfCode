package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// (kdd) TODO: Figure this out!

/*
After trying for many hours to get the input via http request I have come to the realization
that this was won't work. Using these functions the resp comes back blank. Using curl,
the resp come back along the lines of: " Pleasee login, the input is different for every user."
Looking at the page source reveals only the input, so it must be ajax or something.
Since I am unsure of how to continue, I am just going to skip this part, and hard-code
the input to the file.
*/

func getDataFromHTTP(link string, b []byte) error {
	resp, err := http.Get(link)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = resp.Body.Read(b)
	if err != nil {
		log.Println(err)
		return err
	}
	/*
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			return err
		}
		copy(b, []byte(data))
	*/
	return nil
}

func deriveFileNameFromLink(link string) string {
	cwd, _ := os.Getwd()
	short := strings.TrimPrefix(link, "https://adventofcode.com/")
	short = strings.Join(strings.Split(short, "/"), "")
	var fname = strings.Join([]string{cwd, "/", short, ".txt"}, "")
	return fname

}

func writeToFile(fname string, data []byte) error {
	if err := ioutil.WriteFile(fname, data, 0666); err != nil {
		log.Println(err)
		return err
	}
	log.Printf("wrote to file: %s", fname)
	return nil

}
