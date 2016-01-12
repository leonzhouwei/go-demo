package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	path := "/proc/mounts"

	lines1, error1 := readAllLines(path)
	if error1 != nil {
		log.Fatal(error1)
	}
	printLines(lines1)

	fmt.Println("==========")
	lines2, error2 := readAllLinesWithIoutil(path, "\n")
	if error2 != nil {
		log.Fatal(error2)
	}
	printLines(lines2)
}

func printLines(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
	}
}

func readAllLinesWithIoutil(path string, sep string) (lines []string, err error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	str := string(bytes)
	lines = strings.Split(str, sep)
	return
}

func readAllLines(path string) (lines []string, err error) {
	file, err := os.Open(path)

	if err != nil {
		return
	}
	defer file.Close()

	br := bufio.NewReader(file)
	lines = make([]string, 16)
	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}
		if isPrefix {
			err = errors.New("A too long line, seems unexpected")
			return
		}
		str := string(line)
		lines = append(lines, str)
	}
	return
}
