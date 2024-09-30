package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type FileHandler interface {
	handle(path string, info os.FileInfo)
}

type FileHandlerFunc func(path string, info os.FileInfo)

func (f FileHandlerFunc) handle(path string, info os.FileInfo) {
	f(path, info)
}

type LogFileHandler struct {
	parser     *Parser
	fn         FileHandler
	daysToKeep int
}

func (r *LogFileHandler) handle(path string, info os.FileInfo) {
	entry, err := r.parser.ParseString(info.Name())
	if err != nil {
		return
	}
	if entry.IsExpired(r.daysToKeep) {
		r.fn.handle(path, info)
	}
}

func visit(handler FileHandler) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err) // 可以选择如何处理错误
			return nil
		}
		handler.handle(path, info)
		return nil
	}
}

func RemoveLogFiles(dirs []string, filenamePattern string, daysToKeep int, debug bool) {
	if debug {
		fmt.Println("RemoveLogFiles", dirs, filenamePattern, daysToKeep)
	}

	doHandleLogFile(dirs, filenamePattern, daysToKeep, FileHandlerFunc(func(path string, info os.FileInfo) {
		if !info.IsDir() {
			fmt.Println("Removing: ", path)
			os.Remove(path)
		}
	}))
}

func PrintLogFiles(dirs []string, filenamePattern string, daysToKeep int, debug bool) {
	if debug {
		fmt.Println("PrintLogFiles", dirs, filenamePattern, daysToKeep)
	}

	doHandleLogFile(dirs, filenamePattern, daysToKeep, FileHandlerFunc(func(path string, info os.FileInfo) {
		if !info.IsDir() {
			fmt.Println("File: ", path)
		}
	}))
}

func doHandleLogFile(dirs []string, filenamePattern string, daysToKeep int, fn FileHandler) {
	if len(dirs) < 1 {
		return
	}
	handler := &LogFileHandler{parser: NewParser(filenamePattern), fn: fn, daysToKeep: daysToKeep}
	var wg sync.WaitGroup
	wg.Add(len(dirs))
	for _, dir := range dirs {
		go func(dir string) {
			defer wg.Done()
			filepath.Walk(dir, visit(handler))
		}(dir)
	}
	wg.Wait()
}
