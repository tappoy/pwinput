// This package provides a robust way to read password from the terminal.
//
// Features:
//	- Read password from the terminal.
//	- Mute the terminal while inputting the password.
//	- If interrupted while inputting the password, the terminal will be turned back to normal.
//	- If the input is not from the terminal, it will read from the Env.In.
//
// Dependencies:
//	 - github.com/tappoy/env (Env.In)
//	 - golang.org/x/crypto/ssh/terminal
package pwinput

import (
	"errors"
	"io"
	"os"
	"os/signal"
	"syscall"
	"strings"

	"github.com/tappoy/env"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	// The input process is interrupted.
	ErrInterrupted = errors.New("ErrInterrupted")

)

type PasswordInput interface {
	// InputPassword reads a password from the terminal.
	// If DummyPasswordInput is used, it will return the dummy password.
	// Errors:
	//	- ErrInterrupted
	InputPassword() (string, error)
}

// True implementation
type passwordInputImpl struct{}

// Create a new password reader.
func NewPasswordInput() PasswordInput {
	return passwordInputImpl{}
}

// InputPassword reads a password from the terminal.
// If the input is not from the terminal, it will read from the Env.In.
//
// Errors:
//	- ErrInterrupted
func (pw passwordInputImpl) InputPassword() (string, error) {
	if terminal.IsTerminal(int(syscall.Stdin)) {
		return pw.inputFromTerminal()
	} else {
		return pw.inputFromReader(env.In)
	}
}

func (pw passwordInputImpl) inputFromTerminal() (string, error) {
	// caputure the signal of Ctrl+C
	signalChan := make(chan os.Signal, 512)
	signal.Notify(signalChan, os.Interrupt)
	defer signal.Stop(signalChan)

	// create a channel to store the password
	pwChan := make(chan string)
	defer close(pwChan)

	// create a channel to store the error
	errChan := make(chan error)
	defer close(errChan)

	// copy the current terminal state
	currentState, err := terminal.GetState(int(syscall.Stdin))
	if err != nil {
		return "", err
	}

	go func() {
		<-signalChan
		// restore the terminal state after receiving Ctrl+C
		terminal.Restore(int(syscall.Stdin), currentState)
		errChan <- ErrInterrupted
	}()

	go func() {
		pw, err := terminal.ReadPassword(syscall.Stdin)
		if err != nil {
			pwChan <- ""
		} else {
			pwChan <- string(pw)
		}
	}()

	select {
	case pw := <-pwChan:
		return pw, nil
	case err := <-errChan:
		return "", err
	}
}

func (pw passwordInputImpl) inputFromReader(reader io.Reader) (string, error) {
	env.Outf("input password from reader\n")
	b, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(b)), nil
}
