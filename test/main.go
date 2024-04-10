package main

import (
  "github.com/tappoy/pwinput"
  "fmt"
  )


// test
func main() {
  fmt.Println("Enter your password: ")
  password, err := pwinput.ReadPassword()
  if err != nil {
    fmt.Printf("An error occurred: %s\n", err)
  } else {
    fmt.Printf("Your password is: <%s>\n", password)
  }
  fmt.Println("Finished")
}
