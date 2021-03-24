package webcontent

import (
	"bufio"
	"os"
)

// Load each line as a separate splash
func (content *Content) LoadSplashes(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content.splashes = append(content.splashes, scanner.Text())
	}
}

// Unloads any splashes that have previously been loaded
func (content *Content) ClearSplashes(filepath string) {
	content.splashes = nil
}

// Return a randomly picked line from the loaded splashes
func (content *Content) GetRandomSplash() string {
	if len(content.splashes) == 0 {
		panic("Attempted to get a random splash but none are loaded!")
	}

	i := content.random.Intn(len(content.splashes))
	return content.splashes[i]
}
