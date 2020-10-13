package wincd

import (
	"log"

	"github.com/harry1453/go-common-file-dialog/cfd"
)

// use win CFD to select folder
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
