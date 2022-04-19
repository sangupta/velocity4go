/**
 * velocity4go: Velocity template engine for Go
 * https://sangupta.com/projects/velocity4go
 *
 * MIT License.
 * Copyright (c) 2022, Sandeep Gupta.
 *
 * Use of this source code is governed by a MIT style license
 * that can be found in LICENSE file in the code repository.
 */

package utils

import (
	"errors"
)

// ----------------
// Common functions
// ----------------

/**
 * Check if a given rune is a valid alphabet (currently, only English alphabets)
 *
 */
func IsAsciiLetter(char rune) bool {
	if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') {
		return true
	}

	return false
}

/**
 * Check if a rune is a valid ASCII digit (number, 0-9)
 */
func IsAsciiDigit(char rune) bool {
	if char >= '0' && char <= '9' {
		return true
	}

	return false
}

func IsIdChar(char rune) bool {
	if IsAsciiLetter(char) || IsAsciiDigit(char) || char == '-' || char == '_' {
		return true
	}

	return false
}

func AssertRune(actual rune, expected rune) {
	if expected != actual {
		panic(ParseException("Assertion failed: expected=" + string(expected) + ", actual=" + string(actual)))
	}
}

func ParseException(message string) error {
	return errors.New(message)
}

func ParseInt(str string) int {
	panic(ParseException("Invalid integer: " + str))
}
