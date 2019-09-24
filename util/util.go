package util

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func SecondsDuration(n float64) time.Duration {
	d, _ := time.ParseDuration(strconv.FormatFloat(n, 'f', 3, 64) + "s")
	return d
}

func AfterByte(str string, chr byte) string {
	var i int
	for i = len(str) - 1; i >= 0; i-- {
		if chr == str[i] {
			break
		}
	}
	return str[i+1:]
}

func StripWhitespace(str string) string {
	ctr := 0
	f := func(r rune) rune {
		for range Once {
			if !unicode.IsSpace(r) {
				break
			}
			r = ' '
			if ctr == 1 {
				r = -1
				ctr = 0
				break
			}
			ctr++
		}
		return r
	}
	return strings.Map(f, str)
}

func noop(i ...interface{}) interface{} { return i }

func Fail(message string, args ...interface{}) {
	fmt.Printf(message, args...)
	os.Exit(1)
}
