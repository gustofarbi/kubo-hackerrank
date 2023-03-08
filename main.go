package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const welcome = "WELCOME"

func main() {
	if len(os.Args) != 3 {
		fmt.Println("wrong no. of args")
		os.Exit(1)
	}

	_n := os.Args[1]
	_m := os.Args[2]
	n, err := strconv.Atoi(_n)
	check(err)
	m, err := strconv.Atoi(_m)
	check(err)

	if n*3 != m {
		fmt.Println("wrong input")
		os.Exit(1)
	}

	welcomeLen := len(welcome)
	var invert bool
	for i := 0; i < n; i++ {
		if i+1 == n-i {
			dashes := strings.Repeat("-", (m-welcomeLen)/2)
			fmt.Printf("%s%s%s\n", dashes, welcome, dashes)
			invert = true
			continue
		}

		var nMiddle int
		if invert {
			nMiddle = ((n - i - 1) * 2) + 1
		} else {
			nMiddle = i*2 + 1
		}
		middle := strings.Repeat(".|.", nMiddle)
		nDashes := float64((m - len(middle)) / 2)

		dashes := strings.Repeat("-", int(math.Abs(nDashes)))
		fmt.Printf("%s%s%s\n", dashes, middle, dashes)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
