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
			if r+rune(3) > 'z' {
				return (r-'a'+rune(3))%26 + 'a'
			}
			return r + rune(3)
		}, normalizedS)
	return cipheredS
}

// Encode ...
func (c *Shift) Encode(s string) string {
	normalizedS := strings.Map(Normalize, s)
	cipheredS := strings.Map(
		func(r rune) rune {
			switch {
			case c.distance < 0:
				if r+rune(c.distance) < 'a' {
					return 'z' - (r+rune(c.distance))%26
				}
				return r - rune(c.distance*-1)
			default:
				if r+rune(c.distance) > 'z' {
					return (r-'a'+rune(c.distance))%26 + 'a'
				}
				return r + rune(c.distance)
			}
		}, normalizedS)
	return cipheredS
}

// Encode ...
func (c *Vigenere) Encode(s string) string {
	s = strings.Map(Normalize, s)
	fmt.Print(s)
	return ""
}

// Decode ...
func (c *Caesar) Decode(s string) string {
	normalizedS := strings.Map(Normalize, s)
	cipheredS := strings.Map(
		func(r rune) rune {
			if r-rune(3) < 'a' {
				return ('a' - r) + 'z'
			}
			return r - rune(3)
		}, normalizedS)
	return cipheredS
}

// Decode ...
func (c *Shift) Decode(s string) string {
	normalizedS := strings.Map(Normalize, s)
	cipheredS := strings.Map(
		func(r rune) rune {
			if r-rune(c.distance) < 'a' {
				return ('a' - r) + 'z'
			}
			return r - rune(c.distance)
		}, normalizedS)
	return cipheredS
}

// Decode ...
func (c *Vigenere) Decode(s string) string {
	s = strings.Map(Normalize, s)
	fmt.Print(s)
	return ""
}
