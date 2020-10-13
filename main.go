package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/harry1453/go-common-file-dialog/cfd"
)

func main() {
	s := selectFile()
	// fmt.Printf("you picked %s", s)

	csvFile, _ := os.Open(s)
	reader := csv.NewReader(bufio.NewReader(csvFile))

	outFile, err := os.Create(selectFile())
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	csvWriter := csv.NewWriter(csvFile)
	csvWriter.Comma = '|'

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		// write to file
		csvWriter.Write(line)
	}

	csvWriter.Flush()
	outFile.Close()
}

func selectFolder() string {
	pickFolderDialog, err := cfd.NewSelectFolderDialog(cfd.DialogConfig{
		Title: "Pick Folder",
		Role:  "PickFolderExample",
	})
	if err != nil {
		log.Fatal(err)
	}
	if err := pickFolderDialog.Show(); err != nil {
		log.Fatal(err)
	}
	result, err := pickFolderDialog.GetResult()
	if err != nil {
		log.Fatal(err)
	}
	// log.Printf("Chosen folder: %s\n", result)

	return result
}

func selectFile() string {
	openDialog, err := cfd.NewOpenFileDialog(cfd.DialogConfig{
		Title: "Open A File",
		Role:  "OpenFileExample",
		FileFilters: []cfd.FileFilter{
			{
				DisplayName: "CSV Files (*.csv)",
				Pattern:     "*.csv",
			},
			{
				DisplayName: "Excel Files (*.xlsx)",
				Pattern:     "*.xlsx",
			},
			{
				DisplayName: "Text Files (*.txt)",
				Pattern:     "*.txt",
			},
			{
				DisplayName: "All Files (*.*)",
				Pattern:     "*.*",
			},
		},
		SelectedFileFilterIndex: 2,
		FileName:                "file.csv",
		DefaultExtension:        "csv",
	})
	if err != nil {
		log.Fatal(err)
	}

	if err := openDialog.Show(); err != nil {
		log.Fatal(err)
	}

	result, err := openDialog.GetResult()
	if err != nil {
		log.Fatal(err)
	}
	// log.Printf("Chosen file: %s\n", result)

	return result
}
