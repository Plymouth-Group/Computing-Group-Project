// Codinoc IDE
//
// Project by Group 35
// Computing Group Project
//
// Plymouth University UK
// NSBM Green University LK
//
// Password Hashing

package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// BCrypt method is used in password hashing
//
// https://en.wikipedia.org/wiki/Bcrypt

func Hash_Password(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	fmt.Println(fmt.Sprintf("> Password \"%s\" is Hashed: %s", password, string(bytes)))
	return string(bytes)
}

func Check_Hash_Password(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	fmt.Println(fmt.Sprintf("> Password Checked: %s [%s]", password, hash))

	if err == nil {
		return true
	}

	return false
}