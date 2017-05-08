package main

import (
  "bufio"
  "fmt"
  "os"
  "sort"
  "reflect"
)

type IntSlice []int32
func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func deleteStr(s []string, i int) []string {
    s = append(s[:i], s[i+1:]...)
    n := make([]string, len(s))
    copy(n, s)
    return n
}

func deleteInt(s [][]int32, i int) [][]int32 {
    s = append(s[:i], s[i+1:]...)
    n := make([][]int32, len(s))
    copy(n, s)
    return n
}

func main() {
  file, err := os.Open("pla_in.ces")
  if err != nil {
    fmt.Print("error")
  }

  scanner := bufio.NewScanner(file)
  var texts []string
  var sortTexts [][]int32

  for scanner.Scan() {
    texts = append(texts, scanner.Text())
    var chars []int32
    for _, v := range scanner.Text() {
      chars = append(chars, v)
      sort.Sort(IntSlice(chars))
    }
    sortTexts = append(sortTexts, chars)
  }

  for i := 0; i < len(texts); i++ {
    var anagram []string
    anagram = append(anagram, texts[i])
    for j := i+1; j < len(texts); j++ {
      if reflect.DeepEqual(sortTexts[i], sortTexts[j]) {
        anagram = append(anagram, texts[j])
        
        sortTexts = deleteInt(sortTexts, j)
        texts = deleteStr(texts, j)
      }
    }
    if len(anagram) > 1 {
      fmt.Println(anagram)
    }
  }
}
