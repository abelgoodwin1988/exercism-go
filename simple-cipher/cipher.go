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

// Caesar struct to implement the encoding and decoding of a Caeser Cipher,
//	which implements the Cipher interface
type Caesar struct{}

// Shift struct to implement the encoding and decoding of a Shift Cipher,
//	which implements the Cipher interface. The Distance field determine how
//	many alphabet characters to shift.
type Shift struct {
	distance int
}

// Vigenere ...
type Vigenere struct {
	key string
}

// Series of methods for returning new structs
//	which will be used to implement the Cipher interface

// NewCaesar returns Caesar structs which implements the Cipher interface.
//	The Caesar cipher always shifts the alphabetical (english in this case) character three ahead.
//	In the case of extending beyond the last character, the cipher wraps to the beginning
//	of the alphabet
func NewCaesar() *Caesar {
	return &Caesar{}
}

// NewShift returns a new Shift struct which implements the Cipher interface.
//	the Shift cipher shifts n characters based on the defined distance field.
//	Shift accepts negative and positive values, but has bounds based on the
//	English alphabet of 26 and cannot shift 0.
func NewShift(distance int) *Shift {
	if distance <= -26 || distance >= 26 || distance == 0 {
		return nil
	}
	return &Shift{distance: distance}
}

// NewVigenere ...
func NewVigenere(key string) *Vigenere {
	return &Vigenere{key: key}
}

// Encode Caesar returns ciphered string by shifting ahead three characters
//	and wrapping when exceeding the English alphabet character upper-bound of z.
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

// Encode Shift returns ciphered string by shifting the Shift structs set distance
//	field. Wraps when exceeding upper bound or below lower bound.
func (c *Shift) Encode(s string) string {
	normalizedS := strings.Map(Normalize, s)
	cipheredS := strings.Map(
		func(r rune) rune {
			switch {
			case c.distance < 0:
				if r+rune(c.distance) < 'a' {
					return 'z' + 1 + (r + rune(c.distance) - 'a')
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

// Decode Caesar returns deciphered string which used the Encode Caesar method.
//	This is accomplished by going back three english alphabet characters with
//	alphabet wrapping
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

// Decode Shift returns deciphered string which used the Encode Shift method.
//	This is accomplished by going the opposite value defined in the Shift struct.
//	Alphabet wrapping is present.
func (c *Shift) Decode(s string) string {
	normalizedS := strings.Map(Normalize, s)
	cipheredS := strings.Map(
		func(r rune) rune {
			switch {
			case c.distance < 0:
				if r-rune(c.distance) > 'z' {
					return 'a' + (r-rune(c.distance)-'a')%26
				}
			default:
				if r-rune(c.distance) < 'a' {
					return ('a' - r) + 'z'
				}
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
