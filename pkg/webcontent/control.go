package webcontent

// Start the content manager on separate goroutines
func (content *Content) Start(errChan chan<- error) {
	content.errChan = errChan
	content.doneWG.Add(2)
	go content.watchFiles()
	go content.handleFileChanges()
}

// Gracefully shutdown the content manager
func (content *Content) Stop() {
	content.fileWatcher.Close() // TODO: blocks for up to a minute
	content.doneWG.Wait()
}
