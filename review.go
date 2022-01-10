package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// Conduct a brief Code Review of the given code snippets in the context of the previous Task

/*
 * The function storyStats returns four things:
 * the shortest word
 * the longest word
 * the average word length
 * the list of words that are of length equal to the average word rounded up and down, or empty list if such words do not exist.
 */
func storyStats1(str string) (map[string]string, error) {
	var err error
	m := make(map[string]string)
	m["shortest"] = ""
	m["longest"] = ""
	m["average"] = ""

	str = strings.TrimSpace(str)
	if str == "" {
		err = errors.New("can't work empty string")
		return nil, err
	}
	var shortest int
	var longest int
	words := strings.Fields(str)

	for index, element := range words {
		if index == 0 {
			// set intail values
			shortest = len(element)
			longest = len(element)
			m["shortest"] = element
			m["longest"] = element
			continue
		}
		// check longest and shotest words
		if shortest > len(element) {
			m["shortest"] = element
			shortest = len(element)
		}
		if longest < len(element) {
			m["longest"] = element
			longest = len(element)
		}
	}
	m["average"] = fmt.Sprint(averageNumber(str))

	return m, err
}

/*
 * The function testValidity returns `true` if the given strings fits the specs,
 * and `false` otherwise.
 */
func testValidity1(str string) bool {
	patern := `[-]?\d[\d]*[\]?[\d{2}]*?[-]`
	re := regexp.MustCompile(patern)
	return re.MatchString(str)
}
