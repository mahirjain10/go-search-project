package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func recursiveDir(dirPath string, storeFileName *[]string) {
	err := filepath.Walk(dirPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fileExt := filepath.Ext(path)
			fmt.Printf("file path %s : %T", fileExt, fileExt)
			*storeFileName = append(*storeFileName, path)
			// fmt.Println(path, info.Size())
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}
func search(wordToSearch string, fp string) {
	f, err := os.Open(fp)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	r := bufio.NewReader(f)
	b := make([]byte, 1)
	word := ""
	var dirFound string
	found := false
	for {
		n, err := r.Read(b)
		if err == io.EOF {
			fmt.Println("finished reading file")
			break
		}
		if err != nil {
			fmt.Printf("Error %v reading file", err)
			break
		}
		if string(b[n-1]) == " " || string(b[n-1]) == "\n" {
			word = ""
		} else {
			word = word + string(b[n-1])
			// fmt.Println(word)
			if word == wordToSearch {
				found = true
				dirFound = fp
				break
			}
		}
	}
	if found {
		fmt.Printf("word found in %s", dirFound)
	} else {
		fmt.Println("word not found")
	}
}
func main() {

	// Initalializing flags
	dir := flag.String("dir", "", "directory or filename")
	tts := flag.String("tts", "", "text to search")
	ifp := flag.Bool("ifp", true, "is file path or a directory")

	// Parsing flag
	flag.Parse()

	// Access the value of the "dir" flag
	directory := *dir
	textToSearch := *tts
	isFilePath := *ifp

	// Check if the directory or filename flag was provided
	if directory == "" {
		fmt.Println("Please provide a directory or filename using -dir option.")
		return
	}
	// Check if the text to search was provided
	if textToSearch == "" {
		fmt.Println("Please provide a text to search using -tts option.")
	}

	var store []string

	if isFilePath {
		search(textToSearch, directory)
	} else {
		recursiveDir(directory, &store)
		for _, st := range store {
			search(textToSearch, st)
		}
	}

}
