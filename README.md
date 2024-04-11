# Package
`github.com/tappoy/pwinput`

# About
This golang package provides a robust way to read password from the terminal.

# Features
- Read password from the stdin.
- Mute the terminal while inputting the password.
- If interrupted while inputting the password, the terminal will be turned back to normal.

# Functions
- `ReadPassword() (string, error)`

# Errors
- `ErrInterrupted`: The input process is interrupted.

# License
[LGPL-3.0](LICENSE)

# Author
[tappoy](https://github.com/tappoy)
