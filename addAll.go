package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func AddFilesAndFolders(watcher *fsnotify.Watcher, filesOrFolders []string) {
	var err error
	for _, e := range filesOrFolders {
		err = watcher.Add(e)
		if err != nil {
			log.Fatal(err)
		}

	}
}
