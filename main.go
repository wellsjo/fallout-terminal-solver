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
  for _, char := range s {
    m[string(char)] = true
  }
  return len(m)
}

// Main terminal solving logic
func solve(wordList []string) {

  // First guess the word with most unique characters
  sort.Sort(ByUniqueness(wordList))
  var guesses = make(map[string]int)

  for {
    i, guess := NextGuess(wordList, guesses)
    if (i == -1) {
      fmt.Print("There are no valid guesses left.")
      fmt.Println(" Something was entered incorrectly, please try again.")
      break
    }
    fmt.Printf("Enter %v (index %v)\n", guess, i)
    fmt.Print("Likelihood (number): ")
    scanner.Scan()
    likelihood, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
    wordList = append(wordList[:i], wordList[i+1:]...)
    guesses[guess] = likelihood
  }
}

func NextGuess(wordList []string, guesses map[string]int) (int, string) {
  var index int = -1
  var nextGuess string = ""
  if len(guesses) == 0 {
    return 0, wordList[0]
  } else {
    NextGuess: for i, n := range wordList {
      for g := range guesses {
        if (similarities(n, g) != guesses[g]) {
          continue NextGuess
        }
      }
      index, nextGuess = i, n
      break
    }
  }
  return index, nextGuess
}

func similarities(w1, w2 string) int {
  if (len(w1) == 0 || len(w2) == 0) {
    return 0
  }
  s := 0
  if w1[0] == w2[0] {
    s = 1
  }
  return s + similarities(w1[1:], w2[1:])
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
      fmt.Printf("\nYou have entered %v word%v. %v\n", len(wordList), plural, wordList)
      fmt.Println("Enter 'f' to finish, or 'u' to undo.")
    }
  }
  return wordList
}

func main() {
  wordList := getInput()
  solve(wordList)
}
