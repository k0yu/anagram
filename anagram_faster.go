package main

import (
  "bufio"
  "fmt"
  "os"
  "sort"
  "strings"
)

func main()  {
  texts := readFile("pla_in.ces")

  m := map[string]string {}

  for _, v := range texts {
    m[v] = sortChars(v)
  }

  a := List{}
    for k, v := range m {
        e := Entry{k, v}
        a = append(a, e)
    }

    sort.Sort(a)

    current := ""
    previous := ""
    anagrams := []string {}
    for _, entry := range a {
      current = entry.value
      if current == previous {
        anagrams = append(anagrams, entry.name)
      } else {
        if len(anagrams) > 1 {
          fmt.Println(anagrams)
        }
        anagrams = []string {}
        anagrams = append(anagrams, entry.name)
      }
      previous = current
    }
}

func readFile(filePath string) []string {
  file, err := os.Open(filePath)
  if err != nil {
    fmt.Print("error")
  }

  texts := []string {}
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    texts = append(texts, scanner.Text())
  }

  return texts
}

func sortChars(text string) string{
  chars := strings.Split(text, "")
  sort.Strings(chars)
  return strings.Join(chars, "")
}

type Entry struct {
    name  string
    value string
}
type List []Entry

func (l List) Len() int {
    return len(l)
}

func (l List) Swap(i, j int) {
    l[i], l[j] = l[j], l[i]
}

func (l List) Less(i, j int) bool {
    if l[i].value == l[j].value {
        return (l[i].name < l[j].name)
    } else {
        return (l[i].value < l[j].value)
    }
}
