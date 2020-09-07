package aes

import (
	"fmt"
	"os"
)

func Encrypt(fileChannel *chan os.FileInfo, scanCompleteChannel *chan bool) {
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

func doEncrypt(file os.FileInfo) {
	fileCount++
	fmt.Println(file.Name(), " encrypted")
}
