![Github CI/CD](https://img.shields.io/github/workflow/status/booomch/demo-string-specs/Go)
![Repository Top Language](https://img.shields.io/github/languages/top/booomch/demo-string-specs)
![Scrutinizer Code Quality](https://img.shields.io/scrutinizer/quality/g/booomch/demo-string-specs/main)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/booomch/demo-string-specs)
![Codacy Grade](https://img.shields.io/codacy/grade/68c42124afcf456ab0c28e2d6da7e534)
![Github Repository Size](https://img.shields.io/github/repo-size/booomch/demo-string-specs)
![Github Open Issues](https://img.shields.io/github/issues/booomch/demo-string-specs)
![Lines of code](https://img.shields.io/tokei/lines/github/booomch/demo-string-specs)
![GitHub last commit](https://img.shields.io/github/last-commit/booomch/demo-string-specs)
![License](https://img.shields.io/badge/license-Apache%202-blue)
![GitHub contributors](https://img.shields.io/github/contributors/booomch/demo-string-specs)


# Test tasks

**Important**:
* [ ] tasks must be documented as if it was a programming tasks (documented for programmers)
* [ ] task must be implemented with `Go` and  `GIT`
* [ ] task MUST include commits for every logical block of work. It is expected to have frequent commits with meaningful messages, preferably multiple commits per-function.
* [ ] as a solution submit your `Git` repo URL


### String Specs

Given a string that is a sequence of numbers followed by dash followed by text, eg: `23-ab-48-caba-56-haha`
   * The numbers can be assumed to be unsigned integers
   * The strings can be assumed to be ASCII sequences of arbitrary length larger than 0 (empty strings not allowed).

## Task 1 

* Before implementing the needed function, please estimate the difficulty of the functions and report the estimated and the used time to implement the functions in the function comment block.
* Write a function `testValidity` that takes the string as an input, and returns boolean flag `true` if the given string complies with the format, or `false` if the string does not comply
* Write a function `averageNumber` that takes the string, and returns the average number from all the numbers
* Write a function `wholeStory` that takes the string, and returns a text that is composed from all the text words separated by spaces, e.g. `story` called for the string `1-hello-2-world` should return text: `"hello world"
* Write a function `storyStats` that returns four things:
   * the shortest word
   * the longest word
   * the average word length
   * the list (or empty list) of all words from the story that have the length the same as the average length rounded up and down.

# Task 2

* Before implementing the needed function, please estimate the difficulty of the functions and report the estimated and the used time to implement the functions in the function comment block.
* Write a `generate` function, that takes boolean flag and generates random correct strings if the parameter is `true` and random incorrect strings if the flag is `false`.
* Write unit tests for the functions:
   * `testValidity`
   * `averageNumber`
   * `wholeStory`
   * `storyStats`

