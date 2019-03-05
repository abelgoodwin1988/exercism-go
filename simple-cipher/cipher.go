package cipher

import (
	"fmt"
	"strings"
)

// Normalize removes non-alphabet characters and lowercases all characters
//	in a string
func Normalize(r rune) rune {
	if r >= 'A' && r <= 'z' {
		if r >= 'A' && r <= 'Z' {
			return r + 'a' - 'A'
		}
		return r
	}
	return -1
}

// Cipher interface is to be implemented by all -cipher structs
type Cipher interface {
	Encode(string) string
	Decode(string) string
}

// Series of structs used for ciphering which will implement the
//	Cipher interface

// Caesar ...
type Caesar struct{}

// Shift ...
type Shift struct {
	distance int
}

// Vigenere ...
type Vigenere struct {
	key string
}

// Series of methods for returning new structs
//	which will be used to implement the Cipher interface

// NewCaesar ...
func NewCaesar() *Caesar {
	return &Caesar{}
}

// NewShift ...
func NewShift(distance int) *Shift {
	return &Shift{distance: distance}
}

// NewVigenere ...
func NewVigenere(key string) *Vigenere {
	return &Vigenere{key: key}
}

// Encode ...
func (c *Caesar) Encode(s string) string {
	normalizedS := strings.Map(Normalize, s)
	cipheredS := strings.Map(
		func(r rune) rune {
			if r+3 > 'z' {
				return r + 3 - 'a'
			}
			return r + 3
		}, normalizedS)
	return cipheredS
}

// Encode ...
func (c *Shift) Encode(s string) string {
	s = strings.Map(Normalize, s)
	fmt.Print(s)
	return ""
}

// Encode ...
func (c *Vigenere) Encode(s string) string {
	s = strings.Map(Normalize, s)
	fmt.Print(s)
	return ""
}

// Decode ...
func (c *Caesar) Decode(s string) string {
	s = strings.Map(Normalize, s)
	fmt.Print(s)
	return ""
}

// Decode ...
func (c *Shift) Decode(s string) string {
	s = strings.Map(Normalize, s)
	fmt.Print(s)
	return ""
}

// Decode ...
func (c *Vigenere) Decode(s string) string {
	s = strings.Map(Normalize, s)
	fmt.Print(s)
	return ""
}
