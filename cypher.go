package main 

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"bytes"
)

func printMatrix(matrix [][]string, matrix_size int) {
	for i := 0; i < matrix_size; i++ {
		for j := 0; j < matrix_size; j++ { 
			fmt.Printf(matrix[i][j])
		}
		fmt.Printf("\n")
	}
}

func userInput(option string) ([]byte, []byte) {
    reader := bufio.NewReader(os.Stdin)
    var raw_word   string
    var secret_key string
    
    if option == "-encrypt" {
    	fmt.Println("type the word to encrypt: ")
    	} else {
    		fmt.Println("type the word to be decrypted: ")
    	}
    raw_word, _ = reader.ReadString('\n')

    fmt.Println("type the secret key that will be used to encrypt/decrypt: ")
    secret_key, _ = reader.ReadString('\n')

    strings.ToLower(raw_word)		
    strings.ToLower(secret_key)		

    strings.Split(raw_word, "")		// String to String Array
    strings.Split(secret_key, "")	// String to String Array

    // Each char of the String to its ASCII value
    return []byte(raw_word), []byte(secret_key)
}

func encrypt(matrix [][]string, raw_word_charValue byte, secret_key_charValue byte, encrypted_word *bytes.Buffer) {

	// int8 to int
	row		:= int(raw_word_charValue)
    column 	:= int(secret_key_charValue)

    // Minus 97 because array is indexed from 0 to n and a = 97 on ASCII Table
	secret_word := matrix[row-97][column-97]

	// Concatenates to the referenced string variable
	encrypted_word.WriteString(secret_word)
}

func decrypt(matrix [][]string, encrypted_charValue byte, secret_key_charValue byte, decrypted_word *bytes.Buffer) {
	// int8 to int
	var row int
	if encrypted_charValue >= secret_key_charValue {
		row	= int(encrypted_charValue - secret_key_charValue)
	} else {
		row	= int(secret_key_charValue - encrypted_charValue)
		row = 26 - row
	}

    // Minus 97 because array is indexed from 0 to n and a = 97 on ASCII Table
	real_word := matrix[row][0]

	// Concatenates to the referenced string variable
	decrypted_word.WriteString(real_word)
}

func main() {
	//flag from stdin
	option := os.Args[1]

	chars 		:= strings.Split("abcdefghijklmnopqrstuvwxyz", "")
	chars_size 	:= len(chars)

	// Cria a matriz
	matrix := make([][]string, chars_size) 
	for i := 0; i < chars_size; i++ {
		matrix[i] = make([]string, chars_size)
	}

	for i := 0; i < chars_size; i++ {
		// if 0 then starts with a, if 1 then starts with b, etc
		for j := i; j < chars_size; j++ { 
			matrix[i][j-i] = chars[j]
		}
		
		// Adds the leftover/missing words, i.e: started from b-z, adds `a` (was missing)
		for k := 0; k < i; k++ {
			matrix[i][((chars_size-i)+k)] = chars[k]
		}
	}

	printMatrix(matrix, chars_size)

	if option == "-encrypt" {
		fmt.Printf("**** encrypt ****\n")
		// Encrypted Word - Used to concatenate the encrypted word
		var encrypted_word bytes.Buffer

		raw_word, secret_key := userInput(option)

		for i := 0; i < len(raw_word)-1; i++ { // minus 1 cause new line ascii 10 
			encrypt(matrix, raw_word[i], secret_key[i], &encrypted_word)
		}

		fmt.Printf("Encrypted word: %s \n", encrypted_word.String())
	} else if option == "-decrypt" {
		fmt.Printf("**** decrypt ****\n")
		// Decrypted Word - Used to concatenate the encrypted word
		var decrypted_word bytes.Buffer 
		encrypted_word, secret_key := userInput(option)

		// minus 1 cause new line ascii 10 
		for i := 0; i < len(encrypted_word)-1; i++ {
			decrypt(matrix, encrypted_word[i], secret_key[i], &decrypted_word)
		}

		fmt.Printf("Decrypted word: %s \n", decrypted_word.String())
	}
}

