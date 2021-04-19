package webcontent

import (
	"time"
	"log"
)

func (content *Content) watchFiles() {
	defer content.doneWG.Done()
	err := content.fileWatcher.Start(10 * time.Second)
	if err != nil {
		content.errChan <- err
	}
}

func (content *Content) handleFileChanges() {
	defer content.doneWG.Done()
	for {
		select {
			case <-content.fileWatcher.Event:
				log.Println("File change detected, reloading content...")
				content.loadHTMLTemplates()
				content.loadPosts()
				content.loadSplashTexts()
				log.Println("Reloading complete")
			case err := <-content.fileWatcher.Error:
				content.errChan <- err
			case <-content.fileWatcher.Closed:
				return
		}
	}
}
