package pwinput

// Dummy implementation
type dummyPasswordInput struct {
	dummyPassword string
}

var (
	// The dummy password to generate ErrInterrupted.
	Interrupt = "Interrupt"
)

// Return the dummy password.
// Errors:
//   - ErrInterrupted
func (pw dummyPasswordInput) InputPassword() (string, error) {
	if pw.dummyPassword == Interrupt {
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

