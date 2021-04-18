package webcontent

import (
	"time"
	"log"
)

func (content *Content) watchFiles() {
	defer content.doneWG.Done()
	err := content.fileWatcher.Start(1 * time.Minute)
	if err != nil {
		content.errChan <- err
	}
}

func (content *Content) handleFileChanges() {
	defer content.doneWG.Done()
	for {
		select {
			case event := <-content.fileWatcher.Event:
				// TODO: only reload a directory which changed
				log.Println(event)
				content.loadHTMLTemplates()
				content.loadPosts()
				content.loadSplashTexts()
			case err := <-content.fileWatcher.Error:
				content.errChan <- err
			case <-content.fileWatcher.Closed:
				return
		}
	}
}
