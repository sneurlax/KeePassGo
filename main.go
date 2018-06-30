package main

import (
  "fmt"
  "github.com/tobischo/gokeepasslib"
  "os"
)

func main() {
  file, _ := os.Open("test.kdbx")

  db := gokeepasslib.NewDatabase()
  db.Credentials = gokeepasslib.NewPasswordCredentials("test")
  _ = gokeepasslib.NewDecoder(file).Decode(db)

  db.UnlockProtectedEntries()

  entry := db.Content.Root.Groups[0].Entries[0]
  fmt.Println(entry.GetTitle())
  fmt.Println(entry.Get("UserName"))
  fmt.Println(entry.GetPassword())
} 
