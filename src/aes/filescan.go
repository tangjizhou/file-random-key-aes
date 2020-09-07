package aes

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var secondLevelFolders = []string{"src1", "test2", "test3", "test4"}
var scanFileSuffixes = []string{".txt", ".go", ".java", ".js"}
var fileChannel = make(chan os.FileInfo, 100)
var recursiveDepth = 0
var scanComplete = make(chan bool)

func Scan(path string) (*chan os.FileInfo, *chan bool) {
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		if file.IsDir() && contains(secondLevelFolders, file.Name()) {
			go doScan(filepath.Join(path, file.Name()))
		}
	}
	return &fileChannel, &scanComplete
}

func doScan(path string) {
	recursiveDepth++
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		filename := file.Name()
		if file.IsDir() {
			doScan(filepath.Join(path, filename))
		} else if contains(scanFileSuffixes, filepath.Ext(filename)) {
			fmt.Println("file ", filename, " added")
			fileChannel <- file
		}
	}
	recursiveDepth--
	if recursiveDepth == 0 {
		scanComplete <- true
	}
}

func contains(arr []string, target string) bool {
	if len(arr) == 0 {
		return false
	}
	for _, value := range arr {
		if value == target {
			return true
		}
	}
	return false
}
