package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
)

type Configurations struct {
	rootDir  string
	exclude  []string
	logError bool
	delay    int32
}

const PATH = "./.gorun"

func main() {
	configs := GetConfigs()
	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) || event.Has(fsnotify.Remove) {
					fmt.Println("Restarting...")
					Restart()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()
	// Add the paths.
	dirs := getDirs(configs.rootDir)
	watchFiles := ExcludeFiles(configs.exclude, dirs)
	AddFilesAndFolders(watcher, watchFiles)
	// Block main goroutine forever.
	<-make(chan struct{})
}

func getDirs(root string) []string {
	var result []string
	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		result = append(result, e.Name())
	}
	return result
}
