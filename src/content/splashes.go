package content

import (
	"bufio"
	"os"
	"math/rand"
	"time"
)

var splashes []string

// Load each line as a separate splash
func LoadSplashes(filepath string) {
	rand.Seed(time.Now().Unix())

	file, err := os.Open(filepath)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	splashes = nil
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		splashes = append(splashes, scanner.Text())
	}
	if len(splashes) == 0   {
		panic("Could not load any splashes from " + filepath)
	}
}

// Return a randomly picked line from the loaded splashes
func GetRandomSplash() string {
	return splashes[rand.Intn(len(splashes))]
}
