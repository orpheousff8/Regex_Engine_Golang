package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _, _ := reader.ReadLine()
	inputs := strings.Split(string(line), "|")
	regex, text := inputs[0], inputs[1]

	fmt.Println(proceed(regex, text))
}

func proceed(regex string, text string) bool {
	// special rules
	if len(regex) == 0 {
		return true
	}
	if len(regex) == 0 && len(text) == 0 {
		return true
	}
	if len(text) == 0 {
		return false
	}

	textRev := reverse(text)

	// begin and end with
	if regex[0] == '^' && regex[len(regex)-1] == '$' {
		regex = regex[1 : len(regex)-1]
		regexRev := reverse(regex)
		return isMatched(regex, text) && isMatched(regexRev, textRev) || isPrefix(regex, text) && isPrefix(regexRev, textRev)
	}
	// only begin with
	if regex[0] == '^' {
		regex = regex[1:]
		return isMatched(regex, text) || isPrefix(regex, text)
	}
	// only end with
	if regex[len(regex)-1] == '$' {
		regexRev := reverse(regex[:len(regex)-1])
		return isMatched(regexRev, textRev) || isPrefix(regexRev, textRev)
	}
	// normal
	regexRev := reverse(regex)
	return isMatched(regex, text) || isPrefix(regex, text) || isPrefix(regexRev, textRev)
}

func isMatched(regex, text string) bool {
	if len(regex) == 0 {
		return true
	}
	if len(regex) > 0 && len(text) == 0 {
		return true
	}

	if len(regex) > 1 && regex[0] != '\\' {
		switch regex[1] {
		case '?':
			regex2 := strings.Replace(regex, "?", "", 1)
			// 0 repeat OR 1 repeat
			return isMatched(regex[2:], text[1:]) || isMatched(regex2[:], text[:])
		case '*':
			regex2 := strings.Replace(regex, "*", "", 1)
			// 0 repeat OR infinite repeats
			return isMatched(regex[2:], text[1:]) || isEqual(regex2[0:1], text[0:1]) && isMatched(repeatRegex(regex2, len(text)), text)
		case '+':
			regex2 := strings.Replace(regex, "+", "", 1)
			// 1 repeat OR infinite repeats
			return isMatched(regex2[:], text[:]) || isEqual(regex2[0:1], text[0:1]) && isMatched(repeatRegex(regex2, len(text)), text)
		}
	}
	if len(regex) > 1 && regex[0] == '\\' {
		if !isEqual(regex[0:2], text[0:1]) {
			return false
		}
		return isMatched(regex[2:], text[1:])
	}

	if !isEqual(regex[0:1], text[0:1]) {
		return false
	}
	return isMatched(regex[1:], text[1:])
}

func repeatRegex(regex string, n int) string {
	for len(regex) < n {
		regex = regex[0:1] + regex
	}
	return regex
}

func isPrefix(regex, text string) bool {
	if len(text) == 0 {
		return false
	}
	if isEqual(regex, text) {
		return true
	}
	return isPrefix(regex, text[:len(text)-1])
}

func isEqual(r, t string) bool {
	if r == "." && len(t) == 1 {
		return true
	}
	if len(r) > 1 && r[0] == '\\' {
		r = strings.Replace(r, "\\", "", 1)
	}
	if r == t {
		return true
	}
	return false
}

func reverse(s string) (t string) {
	// reverse order except * or +; go with its preceding character
	// no+pe	->  epo+n
	// .*c		->	c.*
	// abcabc	->	cbacba
	// 3\+5		->	5\+3
	for i := len(s) - 1; i >= 0; i-- {
		if i > 0 && (s[i] == '*' || s[i] == '+' || s[i-1] == '\\') {
			t += s[i-1 : i+1]
			i--
			continue
		}
		t += s[i : i+1]
	}
	return
}
