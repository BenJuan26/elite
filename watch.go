package elite

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

// WatchLocationChange will watch logs folder for changes and scan log files for the new location
func WatchLocationChange(f func(location string)) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	go func() {
		var lastSystem string
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) {
					ss, serr := GetStarSystem()
					if serr != nil {
						continue
					}
					if lastSystem != ss {
						f(ss)
						lastSystem = ss
					}

				}
			case _, ok := <-watcher.Errors:
				if !ok {
					return
				}
			}
		}
	}()

	// Add a path.
	err = watcher.Add(defaultLogPath)
	if err != nil {
		return err
	}

	// Block main goroutine forever.
	<-make(chan struct{})
	return nil
}
