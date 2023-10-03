package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	// "fmt"
	// "log"
	// "os"
	// "github.com/fsnotify/fsnotify"
)

type Configurations struct {
	rootDir  string
	exclude  []string
	logError bool
	delay    int32
}

const PATH = "./.gorun"

func main() {
	fmt.Println("hello")
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
					Clean()
					Restart()
					fmt.Print("Restarting...")
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
	dirs := getDirs()
	watchFiles := ExcludeFiles(configs.exclude, dirs)
	AddFilesAndFolders(watcher, watchFiles)
	// Block main goroutine forever.
	<-make(chan struct{})
}

func getDirs() []string {
	var result []string
	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		fmt.Println(e.Name())
		result = append(result, e.Name())
	}
	return result
}
