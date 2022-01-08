package main

import (
	"regexp"
	"strconv"
	"strings"
)

func main() {
}

//Write a function `testValidity` that takes the string as an input,
//and returns boolean flag `true` if the given string complies with
//the format, or `false` if the string does not comply
// Difficulty: Easy
// Estimated time: 10 min
// Elapsed time: 10 min
func testValidity(s string) bool {
	pattern := `[-]?\d[\d]*[\]?[\d{2}]*?[-]`
	re := regexp.MustCompile(pattern)
	return re.MatchString(s)
}

//Write a function `averageNumber` that takes the string,
//and returns the average number from all the numbers
// Difficulty: Easy
// Estimated time: 10 min
// Elapsed time: 15 min
func averageNumber(s string) uint16 {
	pattern := `[0-9]+`
	numbers := regexp.MustCompile(pattern).FindAllString(s, -1)

	if len(numbers) == 0 {
		return 0
	}
	total := 0
	for _, num := range numbers {
		parsedNum, err := strconv.Atoi(num)
		if err != nil {
			continue
		}
		total = total + parsedNum
	}
	return uint16(total / len(numbers))
}

//Write a function `wholeStory` that takes the string,
//and returns a text that is composed from all the text
//words separated by spaces,
//e.g. `story` called for the string `1-hello-2-world`
//should return text: `"hello world"
// Difficulty: Easy
// Estimated time: 10 min
// Elapsed time: 5 min
func wholeStory(s string) string {
	pattern := `[a-zA-Z]+`
	words := regexp.MustCompile(pattern).FindAllString(s, -1)
	return strings.Join(words, " ")
}

//Write a function `storyStats` that returns four things:
// the shortest word
// the longest word
// the average word length
// the list (or empty list) of all words from the story that
// have the length the same as the average length rounded up and down.
// Difficulty: Normal
// Estimated time: 15 min
// Elapsed time:
func storyStats(str string) (*ShortStoryDto, error) {
	return nil, nil
}

type ShortStoryDto struct {
	ShortestWord      string
	LongestWord       string
	AverageWordLength uint16
	AverageWords      []string
}
