package webcontent

import (
	"bufio"
	"os"
	"errors"
)

// Return a randomly picked line from the loaded splash texts
func (content *Content) GetRandomSplashText() string {
	index := content.random.Intn(len(content.splashTexts))
	return content.splashTexts[index]
}

// Load each line from filepath as a separate splash text
func (content *Content) loadSplashTexts() error {

	// Open file
	file, err := os.Open(content.splashTextsFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Read splash texts from file
	content.splashTexts = []string {}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content.splashTexts = append(content.splashTexts , scanner.Text())
	}

	// Don't allow no splash texts to be loaded
	if len(content.splashTexts) == 0 {
		return errors.New("Unable to load any splashes from " + content.splashTextsFilePath)
	}

	return nil
}
