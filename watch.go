package elite

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

// WatchLocationChange will watch logs folder for changes and scan log files for the new location
func WatchLocationChange(f func(system, station string)) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	go func() {
		var lastSystem string
		var lastStation string
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) {
					system, serr := GetStarSystem()
					if serr != nil {
						continue
					}

					station, serr := GetCurrentStation()
					if serr != nil {
						continue
					}

					if lastSystem != system || lastStation != station {
						f(system, station)
						lastSystem = system
						lastStation = station
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
