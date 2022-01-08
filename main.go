package main

func main() {
}

//Write a function `testValidity` that takes the string as an input,
//and returns boolean flag `true` if the given string complies with
//the format, or `false` if the string does not comply
// Difficulty: Easy
// Estimated time: 10 min
// Elapsed time:
func testValidity(s string) bool {
	return false
}

//Write a function `averageNumber` that takes the string,
//and returns the average number from all the numbers
// Difficulty: Easy
// Estimated time: 10 min
// Elapsed time:
func averageNumber(s string) uint16 {
	return 0
}

//Write a function `wholeStory` that takes the string,
//and returns a text that is composed from all the text
//words separated by spaces,
//e.g. `story` called for the string `1-hello-2-world`
//should return text: `"hello world"
// Difficulty: Easy
// Estimated time: 10 min
// Elapsed time:
func wholeStory(s string) string {
	return ""
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
