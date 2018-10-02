package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "strings"
)

func main() {
  filemap := make(map[string]map[string]int)
  counts := make(map[string]int)

  for _, filename := range os.Args[1:] {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
      fmt.Fprintf(os.Stderr, "dup3: %v", err)
      continue
    }

    for _, line := range strings.Split(string(data), "\n") {
      counts[line]++
      if filemap[line] == nil {
        filemap[line] = make(map[string]int)
      }
      filemap[line][filename]++
    }
  }

  for line, n := range counts {
    if n > 1 {

      fmt.Printf("%d\t%s ", n, line)
      for filename, i := range filemap[line] {
        if i > 0 {
          fmt.Printf("%s ", filename)
        }
      }
      fmt.Println()
    }
  }
}
