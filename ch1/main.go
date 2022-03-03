package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count bytes")
	flag.Parse()
	fmt.Println(count(os.Stdin, *lines, *bytes))
}

func count(r io.Reader, countLines bool, countBytes bool) int {
	scanner := bufio.NewScanner(r)

	if countBytes {
		for scanner.Scan() {
			b := len(scanner.Bytes())
			return b
		}
	}
	if !countLines {
		scanner.Split(bufio.ScanWords)
	}
	wc := 0
	for scanner.Scan() {
		wc++
	}
	return wc
}
