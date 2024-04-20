package main

import (
  "github.com/tappoy/pwinput"
  "fmt"
  )


// test
func main() {
  fmt.Println("Enter your password: ")
  pwi := pwinput.NewPasswordInput()
  password, err := pwi.InputPassword()
  if err != nil {
    fmt.Printf("An error occurred: %s\n", err)
  } else {
    fmt.Printf("Your password is: <%s>\n", password)
  }
  fmt.Println("Finished")
}
