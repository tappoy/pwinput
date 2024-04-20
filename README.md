# Package
`github.com/tappoy/pwinput`

# About
This golang package provides a robust way to read password from the terminal.

# Features
- Read password from the stdin.
- Mute the terminal while inputting the password.
- If interrupted while inputting the password, the terminal will be turned back to normal.

# Interfaces
```go
type PasswordInput interface {
  InputPassword() (string, error)
}
```

# Functions
- `NewPasswordInput() PasswordInput`: Create a new password reader.
- `NewDummyPasswordInput(dummyPassword string) PasswordInput`: Create a dummy password reader for testing. It will return the dummy password when calling `Inputssword()`.

# Errors
- `ErrInterrupted`: The input process is interrupted.

# License
[LGPL-3.0](LICENSE)

# Author
[tappoy](https://github.com/tappoy)
