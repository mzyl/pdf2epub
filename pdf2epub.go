package main

import (
  "net/http"
  "os"
)

func OpenFile() (string, error) {
  f, err := os.Open("DADOES.pdf")
  if err != nil {
    panic(err)
  }
  defer f.Close()
  return f, err
}

func GetFileContentType(out *os.File) (string, error) {
  
  // only the first 512 bytes are use to sniff the content type
  buffer := make([]byte, 512)

  _, err := out.Read(buffer)
  if err != nil {
    return "", err
  }

  // use the net/http package's handy DetectContentType function. Always returns a valid
  // content-type by returning "application/octet-stream" if no others seemed to match
  contentType := http.DetectContentType(buffer)

  return contentType, nil
}
