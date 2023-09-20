package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"time"
)

const (
	letterBytes  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialBytes = "!@#$%^&*()_+-=[]{}\\|;':\",.<>/?`~"
	numBytes     = "0123456789"
)

func generatePassword(length int, useLetters, useSpecialChar, useNums bool) string {
	bytePassword := make([]byte, length)
	for i := range bytePassword {
		if useLetters {
			bytePassword[i] = letterBytes[rand.Intn(len(letterBytes))]
		} else {
			if useSpecialChar {
				bytePassword[i] = specialBytes[rand.Intn(len(specialBytes))]
			} else {
				bytePassword[i] = numBytes[rand.Intn(len(numBytes))]
			}
		}
	}
	return string(bytePassword)
}

func main() {
	log.SetOutput(os.Stdout)
	rand.New(rand.NewSource(time.Now().Unix()))
	useLetters := flag.Bool("l", true, "this switch helps us to generate strong enough password")
	useSpecialChars := flag.Bool("s", true, "this switch helps us to generate strong enough password")
	useNumbers := flag.Bool("n", true, "this switch helps us to generate strong enough password")
	flag.Parse()

	log.Println(generatePassword(45, *useLetters, *useSpecialChars, *useNumbers))
}
