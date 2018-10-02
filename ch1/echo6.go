package main

import (
  "fmt"
  "os"
  "strings"
  "time"
)

func main() {
  //Test 1: for loop benchmark
  start := time.Now()
  s, sep := "", ""
  for _, arg := range os.Args[1:] {
    s += sep + arg
    sep = " "
  }
  fmt.Println(s)
  secs  := time.Since(start).Seconds()
  fmt.Printf("benchmark 1: %.10f seconds\n", secs)

  //Test 2: Another for loop benchmark
  start = time.Now()
  s, sep = "", ""
  for i := 1; i < len(os.Args); i++ {
    s += sep + os.Args[i]
    sep = " "
  }
  fmt.Println(s)
  secs = time.Since(start).Seconds()
  fmt.Printf("benchmark 2: %.10f seconds\n", secs)

  //string join benchmark
  start = time.Now()
  fmt.Println(strings.Join(os.Args[1:], " "))
  secs = time.Since(start).Seconds()

  fmt.Printf("benchmark 3: %.10f seconds\n", secs)
}
