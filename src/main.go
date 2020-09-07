package main

import (
	"./aes"
	"flag"
	"fmt"
	"os"
)

var (
	encrypt bool
	decrypt bool
	path    string
	aesKey  string
	confirm bool
)

func main() {

	defer func() {
		if e := recover(); e != nil {
			fmt.Println("process failed.", e)
			os.Exit(1)
		}
	}()

	flag.BoolVar(&encrypt, "e", false, "encryption mode")
	flag.BoolVar(&decrypt, "d", false, "decryption mode")
	flag.BoolVar(&confirm, "y", false, "confirm")
	flag.StringVar(&path, "p", "", "file path")
	flag.StringVar(&aesKey, "key", "", "decrypt key")
	flag.Parse()
	validate()

	fileChannel, scanCompleteChannel := aes.Scan(path)
	defer close(*fileChannel)
	defer close(*scanCompleteChannel)

	if encrypt {
		aes.Encrypt(fileChannel, scanCompleteChannel)
	} else if decrypt {
		aes.Decrypt(fileChannel, scanCompleteChannel, "")
	}

}

func validate() {
	if !(encrypt || decrypt) {
		panic("process mode required,encrypt or decrypt")
	}

	if decrypt && aesKey == "" {
		panic("decrypt key required")
	}

	if path == "" {
		panic("file path required")
	}

	if confirm {
		return
	}
	// ask for confirm
	fmt.Print("confirm encrypt file in the folder:", path, ",yes or no ?")
	var answer string
	_, err := fmt.Scanf("%s", &answer)
	if err != nil {
		panic(err)
	}
	if answer != "yes" && answer != "y" {
		panic("process canceled")
	}

}
