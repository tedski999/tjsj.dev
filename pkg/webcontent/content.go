package webcontent

import (
	"html/template"
	"math/rand"
)

type Content struct {
	templateDirPath, postsDirPath, splashTextsFilePath string
	htmlTemplates *template.Template
	posts map[string]Post
	splashTexts []string
	random *rand.Rand
}

func Create(templateDirPath, postsDirPath, splashTextsFilePath string) (*Content, error) {

	// Setup content manager
	content := &Content {
		templateDirPath: templateDirPath,
		postsDirPath: postsDirPath,
		splashTextsFilePath: splashTextsFilePath,
		random: rand.New(rand.NewSource(0)),
	}

	// Load content
	if err := content.loadHTMLTemplates(); err != nil { return nil, err }
	if err := content.loadPosts(); err != nil { return nil, err }
	if err := content.loadSplashTexts(); err != nil { return nil, err }

	return content, nil
}
