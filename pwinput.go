// This package provides a robust way to read password from the terminal.
//
//   - Read password from the stdin.
//   - Mute the terminal while inputting the password.
//   - If interrupted while inputting the password, the terminal will be turned back to normal.
package pwinput

import (
	"errors"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

var (
	// The input process is interrupted.
	ErrInterrupted = errors.New("ErrInterrupted")

	// The dummy password to generate ErrInterrupted.
	Interrput = "Interrupt"
)

type PasswordInput interface {
	// InputPassword reads a password from the terminal.
	// If DummyPasswordInput is used, it will return the dummy password.
	// Errors:
	//   - ErrInterrupted
	InputPassword() (string, error)
}

// Dummy implementation
type dummyPasswordInput struct {
	dummyPassword string
}

// Return the dummy password.
// Errors:
//   - ErrInterrupted
func (pw dummyPasswordInput) InputPassword() (string, error) {
	if pw.dummyPassword == "Interrupt" {
		return "", ErrInterrupted
	} else {
		return pw.dummyPassword, nil
	}
}

// Create a dummy password reader for testing or some other purposes.
// It will return the dummy password when calling `Inputssword()`.
//
// If dummyPassword is "Interrupt" or pwinput.Interrupt,
// it will return ErrInterrupted when calling `InputPassword()`
func NewDummyPasswordInput(dummyPassword string) PasswordInput {
	return dummyPasswordInput{dummyPassword: dummyPassword}
}

// True implementation
type passwordInputImpl struct{}

// InputPassword reads a password from the terminal.
func (pw passwordInputImpl) InputPassword() (string, error) {
	// caputure the signal of Ctrl+C
	signalChan := make(chan os.Signal)
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

// Create a new password reader.
func NewPasswordInput() PasswordInput {
	return passwordInputImpl{}
}
