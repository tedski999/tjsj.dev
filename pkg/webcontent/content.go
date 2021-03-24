package webcontent

import (
	"html/template"
	"math/rand"
	"time"
)

type Content struct {
	htmlTemplates *template.Template
	splashes []string
	random *rand.Rand
}

func Create() *Content {
	random := rand.New(rand.NewSource(time.Now().Unix()))

	content := &Content {
		htmlTemplates: nil,
		random: random,
	}

	return content
}
