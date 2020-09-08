package file

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func Encrypt(fileChannel *chan PathedFile, scanCompleteChannel *chan bool) {
	if len(*fileChannel) == 0 {
		fmt.Println("no file found to encrypt.exit.")
		return
	}
	for true {
		select {
		case file := <-*fileChannel:
			doEncrypt(file)
		case <-*scanCompleteChannel:
			var length = len(*fileChannel)
			for i := 0; i < length; i++ {
				doEncrypt(<-*fileChannel)
			}
			goto end
		}
	}
end:
	fmt.Println("encrypt complete,total: ", fileCount)
}

func doEncrypt(file PathedFile) {
	fileCount++
	fmt.Println(file.info.Name(), " encrypted")

}

func paddingText(str []byte, blockSize int) []byte {
	paddingCount := blockSize - len(str)%blockSize
	paddingStr := bytes.Repeat([]byte{byte(paddingCount)}, paddingCount)
	newPaddingStr := append(str, paddingStr...)
	return newPaddingStr
}

func encrypt(src, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(nil)
		return nil
	}
	src = paddingText(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	blockMode.CryptBlocks(src, src)
	return src
}
