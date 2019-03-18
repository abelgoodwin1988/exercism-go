package cipher

import (
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

// Vigenere struct to implement encoding and decoding of a Vigenere Cipher,
//	which implements the Ciper interface. The key field uses a string as the
//	cipher.
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

// NewVigenere accepts a string to be used as a cipher. The characters represent their
//	index in the corresponding alphabet; so a = index 0, d = index 3. Applying the cipher
//	by matching cipher index to subject string index we get an encoded string by shifting
//	the subect string by the ciphers corresponding index as the distance.
func NewVigenere(key string) *Vigenere {
	// Invalidate invalid keys with nil return
	if len(key) <= 2 {
		return nil
	}
	for _, r := range key {
		if r < 'a' || r > 'z' {
			return nil
		}
	}
	return &Vigenere{key: key}
}

// Encode Caesar returns ciphered string by shifting ahead three characters
//	and wrapping when exceeding the English alphabet character upper-bound of z.
func (c *Caesar) Encode(s string) string {
	normalizedS := strings.Map(Normalize, s)
	encodedS := strings.Map(
		func(r rune) rune {
			if r+rune(3) > 'z' {
				return (r-'a'+rune(3))%26 + 'a'
			}
			return r + rune(3)
		}, normalizedS)
	return encodedS
}

// Encode Shift returns ciphered string by shifting the Shift structs set distance
//	field. Wraps when exceeding upper bound or below lower bound.
func (c *Shift) Encode(s string) string {
	normalizedS := strings.Map(Normalize, s)
	encodedS := strings.Map(
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
	return encodedS
}

// Encode Vigenere is a method on the Vigenere struct which utilizes
//	the struct field "key" as a cipher, as described in NewVigenere,
//	and returns a string encoded using the Vigenere method.
func (c *Vigenere) Encode(s string) string {
	normalizedS := strings.Map(Normalize, s)
	// Apply the same-string-index key-alphabet-distance as a shift cipher to the
	//	subject string
	i := 0
	encodedS := strings.Map(
		func(r rune) rune {
			defer func() { i++ }()
			nextPos := r + (rune(c.key[i%len(c.key)]) - 'a')
			if nextPos < 'a' {
				return 'z' - ('a' - nextPos)
			} else if nextPos > 'z' {
				return 'a' + nextPos - 'z' - 1
			}
			return nextPos
		}, normalizedS)
	return encodedS
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

// Decode Vigenere returns deciphered string which used the Encode Vigenere method.
func (c *Vigenere) Decode(s string) string {
	normalizedS := strings.Map(Normalize, s)
	// Apply the same-string-index key-alphabet-distance as a shift cipher to the
	//	subject string
	i := 0
	decodedS := strings.Map(
		func(r rune) rune {
			defer func() { i++ }()
			nextPos := r - (rune(c.key[i%len(c.key)]) - 'a')
			if nextPos < 'a' {
				return 'z' - ('a' - nextPos - 1)
			} else if nextPos > 'z' {
				return 'a' + nextPos - 'z' - 1
			}
			return nextPos
		}, normalizedS)
	return decodedS
}
