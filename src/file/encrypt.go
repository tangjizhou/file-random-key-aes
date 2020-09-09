package file

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
)

var encryptKey = generateKey()
var encryptCount = 0

func Encrypt(fileChannel *chan PathedFile, scanCompleteChannel *chan bool) {
	if len(*fileChannel) == 0 {
		fmt.Println("no file found to encryptAesCBC.exit.")
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
	fmt.Println("encryptAesCBC complete,total: ", decryptCount)
	if encryptCount != 0 {
		fmt.Println("key:", encryptKey)
	}
}

func doEncrypt(file PathedFile) {
	plainBytes, err := ioutil.ReadFile(file.path)
	if err != nil {
		panic(err)
	}
	encryptedBytes := encryptAesCBC(plainBytes, []byte(encryptKey))
	err = ioutil.WriteFile(file.path, encryptedBytes, file.info.Mode())
	if err != nil {
		panic("encryptAesCBC file[ " + file.path + " ] error")
	}
	encryptCount++
	fmt.Println(file.info.Name(), " encrypted")
}

func paddingText(str []byte, blockSize int) []byte {
	paddingCount := blockSize - len(str)%blockSize
	paddingStr := bytes.Repeat([]byte{byte(paddingCount)}, paddingCount)
	newPaddingStr := append(str, paddingStr...)
	return newPaddingStr
}

func encryptAesCBC(src, key []byte) []byte {
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

func generateKey() string {
	key := uuid.Must(uuid.NewV4())
	return string(key[:aes.BlockSize*2])
}
