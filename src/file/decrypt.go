package file

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
)

var fileCount = 0

func Decrypt(fileChannel *chan os.FileInfo, scanCompleteChannel *chan bool, aesKey string) {
	if len(*fileChannel) == 0 {
		fmt.Println("no file found to decrypt.exit.")
		return
	}
	for true {
		select {
		case file := <-*fileChannel:
			doDecrypt(file, aesKey)
		case <-*scanCompleteChannel:
			var length = len(*fileChannel)
			for i := 0; i < length; i++ {
				doDecrypt(<-*fileChannel, aesKey)
			}
			goto end
		}
	}
end:
	fmt.Println("decrypt complete,total: ", fileCount)
}

func doDecrypt(file os.FileInfo, aesKey string) {
	fileCount++
	fmt.Println(file.Name(), " decrypted")
}

func decrypt(src, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(nil)
		return nil
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	blockMode.CryptBlocks(src, src)
	src = unPaddingText(src)
	return src
}

func unPaddingText(str []byte) []byte {
	n := len(str)
	count := int(str[n-1])
	newPaddingText := str[:n-count]
	return newPaddingText
}
