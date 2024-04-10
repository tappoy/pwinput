package pwinput

import (
	"errors"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func ReadPassword() (string, error) {
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
		errChan <- errors.New("ErrInterrupted")
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
