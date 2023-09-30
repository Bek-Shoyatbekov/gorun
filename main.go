package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

// "github.com/fsnotify/fsnotify"

// "github.com/fsnotify/fsnotify"
// "log"

type Configurations struct {
	rootDir  string
	exclude  []string
	logError bool
	delay    int32
}

const PATH = "./.gorun"

func main() {
	// configs := GetConfigs()
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
	// TODO add all paths to watcher
	err = watcher.Add("./file.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Block main goroutine forever.
	<-make(chan struct{})
}
