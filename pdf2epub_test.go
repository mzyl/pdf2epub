package main

import (
  "testing"
)

func TestPDF(t *testing.T) {
  pdf, _ := OpenFile()
  got, _ := GetFileContentType(pdf)
  want := "application/pdf"

  if got != want {
    t.Errorf("got type %q want type %q", got, want)
  }
}
