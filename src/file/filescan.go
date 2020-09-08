package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sync/atomic"
)

var scanFolders = []string{"src1", "test2", "test3", "test4"}
var scanFileSuffixes = []string{".txt", ".go", ".java", ".js"}
var fileChannel = make(chan PathedFile, 100)
var scanCompleteChannel = make(chan bool, 1)

type PathedFile struct {
	info os.FileInfo
	path string
}

func Scan(path string) (*chan PathedFile, *chan bool) {
	files, _ := ioutil.ReadDir(path)

	// 统计匹配的第二层文件夹
	matchedDirs := make([]os.FileInfo, 0)
	for _, file := range files {
		if isDirMatched(file) {
			matchedDirs = append(matchedDirs, file)
		}
	}

	var pendingCount = int32(len(matchedDirs))
	if pendingCount == 0 {
		scanCompleteChannel <- true
	}

	for _, file := range matchedDirs {
		if isDirMatched(file) {
			go doScan(filepath.Join(path, file.Name()), &pendingCount)
		}
	}
	return &fileChannel, &scanCompleteChannel
}

func isDirMatched(file os.FileInfo) bool {
	return file.IsDir() && contains(scanFolders, file.Name())
}

func doScan(rootPath string, count *int32) {
	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return filepath.SkipDir
		}
		if contains(scanFileSuffixes, filepath.Ext(info.Name())) {
			fileChannel <- PathedFile{info, path}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	atomic.AddInt32(count, -1)
	if *count == 0 {
		scanCompleteChannel <- true
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
