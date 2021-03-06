package main

import (
  "fmt"
  "os"
  "path/filepath"
  "strings"
)

func main() {
  path := os.Args[1]
  relative := os.Args[2]
  relPath, err := filepath.Rel(path, relative)
  if err != nil {
    panic(err)
  }
  if !strings.HasPrefix(relPath, "../") {
    relPath = "./" + relPath
  }
  // Remove file extensions
  relPath = strings.TrimSuffix(relPath, filepath.Ext(relPath))
  // Print final result without newline
  fmt.Print("'" + relPath + "'")
}
