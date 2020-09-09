package main

import (
	"./file"
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

	flag.Usage = func() {
		fmt.Println("Note: you must have [rwx] permission of target dir before your start")
		flag.PrintDefaults()
		fmt.Println("example for encrypt: file-aes [-e] -p <path> [-y]")
		fmt.Println("example for decrypt: file-aes -d -key <key> -path <path> [-y]")
	}

	flag.BoolVar(&encrypt, "e", true, "encryption mode")
	flag.BoolVar(&decrypt, "d", false, "decryption mode")
	flag.BoolVar(&confirm, "y", false, "confirm")
	flag.StringVar(&path, "p", "", "file path")
	flag.StringVar(&aesKey, "key", "", "decrypt key")
	flag.Parse()

	validate()
	askConfirm()

	fileChannel, scanCompleteChannel := file.Scan(path)
	defer close(*fileChannel)
	defer close(*scanCompleteChannel)

	if decrypt {
		file.Decrypt(fileChannel, scanCompleteChannel, "")
	} else if encrypt {
		file.Encrypt(fileChannel, scanCompleteChannel)
	} else {
		fmt.Println("mode not exists,exit")
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
}

func askConfirm() {
	if confirm {
		return
	}
	fmt.Print("is the path confirmed [", path, "],yes or no: ")
	var answer string
	_, err := fmt.Scanf("%s", &answer)
	if err != nil {
		panic(err)
	}
	if answer != "yes" && answer != "y" {
		panic("operation canceled")
	}
}
