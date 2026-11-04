package models

import (
	"fmt"
	"net/mail"
	"strings"
	"unicode/utf8"
)

const (
	minPasswordLen   = 8
	maxPasswordBytes = 72
	maxNameRunes     = 100
)

type RegisterInput struct {
	Email    string
	Password string
	Name     string
}

func (in *RegisterInput) Validate() error {
	in.Email = NormalizeEmail(in.Email)
	in.Name = strings.TrimSpace(in.Name)

	if addr, err := mail.ParseAddress(in.Email); err != nil || addr.Address != in.Email {
		return fmt.Errorf("%w: invalid email", ErrInvalidInput)
	}
	if n := utf8.RuneCountInString(in.Name); n == 0 || n > maxNameRunes {
		return fmt.Errorf("%w: name must be 1 to %d characters", ErrInvalidInput, maxNameRunes)
	}
	if strings.ContainsRune(in.Name, '\x00') {
		return fmt.Errorf("%w: name must not contain NUL", ErrInvalidInput)
	}
	if utf8.RuneCountInString(in.Password) < minPasswordLen {
		return fmt.Errorf("%w: password must be at least %d characters", ErrInvalidInput, minPasswordLen)
	}
	if len(in.Password) > maxPasswordBytes {
		return fmt.Errorf("%w: password must be at most %d bytes", ErrInvalidInput, maxPasswordBytes)
	}
	return nil
}
