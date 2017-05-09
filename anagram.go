package main

import (
  "bufio"
  "fmt"
  "os"
  "sort"
  "strings"
)


func deleteStr(s []string, i int) []string {
    s = append(s[:i], s[i+1:]...)
    n := make([]string, len(s))
    copy(n, s)
    return n
}

func sortChars(text string) string{
  chars := strings.Split(text, "")
  sort.Strings(chars)
  return strings.Join(chars, "")
}

func main() {
  file, err := os.Open("pla_in.ces")
  if err != nil {
    fmt.Print("error")
  }

  scanner := bufio.NewScanner(file)
  var texts []string
  var sortTexts []string

  for scanner.Scan() {
    text := scanner.Text()
    texts = append(texts, text)
    sortTexts = append(sortTexts, sortChars(text))
  }

  for i := 0; i < len(texts); i++ {
    var anagram []string
    anagram = append(anagram, texts[i])
    for j := i+1; j < len(texts); j++ {
      if sortTexts[i] == sortTexts[j] {
        anagram = append(anagram, texts[j])

        sortTexts = deleteStr(sortTexts, j)
        texts = deleteStr(texts, j)
      }
    }
    if len(anagram) > 1 {
      fmt.Println(anagram)
    }
  }
}
