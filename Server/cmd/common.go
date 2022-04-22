// Codinoc IDE
//
// Project by Group 35
// Computing Group Project
//
// Plymouth University UK
// NSBM Green University LK
//
// Common Functions

package main

import (
	"golang.org/x/crypto/bcrypt"
)

// BCrypt method is used in password hashing
//
// https://en.wikipedia.org/wiki/Bcrypt

func Hash_Password(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes)
}

func Check_Password_Hash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

// Check Error

func CheckError(err error) {
	if err != nil {
			panic(err)
	}
}