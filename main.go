package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/harry1453/go-common-file-dialog/cfd"
)

func main() {
	s := selectFile()
	// fmt.Printf("you picked %s", s)

	// input/reader
	fi, _ := os.Open(s)
	r := csv.NewReader(fi)

	// Iterate through the records
	for {
		// Read each record from csv
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Reader value: %s\n", rec[0])
	}

	// output/writer
	// fo, err := os.Create(selectFile())
	// if err != nil {
	// 	log.Fatalf("failed creating file: %s", err)
	// }

	// w := csv.NewWriter(fo)
	// w.Comma = '|'

	/*
		for {
			line, error := reader.Read()
			if error == io.EOF {
				break
			} else if error != nil {
				log.Fatal(error)
			}
			fmt.Printf("Length of line is: %s", strconv.Itoa(len(line)))
			fmt.Println
			// write to file
			csvWriter.Write(line)
		}*/

	// w.Flush()
	// fo.Close()
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
