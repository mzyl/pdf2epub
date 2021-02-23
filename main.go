package main

import (
  "fmt"
  "os"
)

func main() {
  
  // open file
  f, err := os.Open("DADOES.pdf")
  if err != nil {
    panic(err)
  }
  defer f.Close()

  // get the content
  pdf, _ := GetFileContentType(f)

  fmt.Println("Content Type: " + pdf)
}
