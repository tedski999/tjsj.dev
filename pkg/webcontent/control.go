package webcontent

// Start the content manager on separate goroutines
func (content *Content) Start(errChan chan<- error) {
	content.errChan = errChan
	content.doneWG.Add(2)
	go content.watchFiles()
	content.fileWatcher.Wait()
	go content.handleFileChanges()
}

// Gracefully shutdown the content manager
func (content *Content) Stop() {
	content.fileWatcher.Close()
	content.doneWG.Wait()
}
