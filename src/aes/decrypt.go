package aes

import (
	"fmt"
	"os"
)

var fileCount = 0

func Decrypt(fileChannel *chan os.FileInfo, scanCompleteChannel *chan bool, aesKey string) {
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
