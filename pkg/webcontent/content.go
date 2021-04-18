package webcontent

import (
	"sync"
	"html/template"
	"math/rand"
	"github.com/radovskyb/watcher"
)

type Content struct {
	templateDirPath, postsDirPath, splashTextsFilePath string
	htmlTemplates *template.Template
	posts map[string]Post
	splashTexts []string
	fileWatcher *watcher.Watcher
	random *rand.Rand
	doneWG sync.WaitGroup
	errChan chan<- error
}

func Create(templateDirPath, postsDirPath, splashTextsFilePath string) (*Content, error) {

	// Setup content manager
	fileWatcher := watcher.New()
	content := &Content {
		templateDirPath: templateDirPath,
		postsDirPath: postsDirPath,
		splashTextsFilePath: splashTextsFilePath,
		fileWatcher: fileWatcher,
		random: rand.New(rand.NewSource(0)),
	}

	// Load content
	if err := content.loadHTMLTemplates(); err != nil { return nil, err }
	if err := content.loadPosts(); err != nil { return nil, err }
	if err := content.loadSplashTexts(); err != nil { return nil, err }

	// Add file watchers
    if err := fileWatcher.AddRecursive(content.templateDirPath); err != nil { return nil, err }
    if err := fileWatcher.AddRecursive(content.postsDirPath); err != nil { return nil, err }
    if err := fileWatcher.AddRecursive(content.splashTextsFilePath); err != nil { return nil, err }

	return content, nil
}
