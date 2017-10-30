package main

import (
  "os"
  "fmt"
  "sort"
  "bufio"
  "strings"
  "strconv"
)

// var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
var scanner = bufio.NewScanner(os.Stdin)

// Type to sort by string uniqueness
type ByUniqueness []string

func (s ByUniqueness) Len() int {
  return len(s)
}

func (s ByUniqueness) Swap(i, j int) {
  s[i], s[j] = s[j], s[i]
}

func (s ByUniqueness) Less(i, j int) bool {
  return uniqueLetters(s[i]) > uniqueLetters(s[j])
}

// returns the # of unique letters in a string
func uniqueLetters(s string) int {
  var m = make(map[string]bool)
  for _, char := range(s) {
    m[string(char)] = true
  }
  return len(m)
}

// Main terminal solving logic
func solve(wordList []string) {
  sort.Sort(ByUniqueness(wordList))
  for {
    fmt.Printf("guess %v\n", wordList[0])
    fmt.Println("Enter the # of matches (likelihood)")
    scanner.Scan()
    matches, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
  }
}

// Input loop to create word list
func getInput() []string {
  var wordList []string
  fmt.Println("Enter words in terminal...")

  MainLoop: for scanner.Scan() {
    word := strings.TrimSpace(scanner.Text())
    switch word {
    case "":
      fmt.Println("Please enter a valid word.")
    case "f":
      break MainLoop
    case "u":
      wordList = wordList[:len(wordList)-1]
      fmt.Printf("Removed last word. %v\n", wordList)
    default:
      wordList = append(wordList, word)
      plural := ""
      if (len(wordList) > 1) {
        plural = "s"
      }
      fmt.Printf("You have entered %v word%v. %v\n", len(wordList), plural, wordList)
      fmt.Println("Enter 'f' to finish, or 'u' to undo.")
    }
  }
  return wordList
}

func main() {
  wordList := getInput()
  solve(wordList)
}
