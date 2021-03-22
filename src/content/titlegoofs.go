package content

import (
	"bufio"
	"os"
	"math/rand"
	"time"
)

var titleGoofs []string

// Load each line as a separate goof
func LoadTitleGoofs(filepath string) {
	rand.Seed(time.Now().Unix())

	file, err := os.Open(filepath)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	titleGoofs = nil
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		titleGoofs = append(titleGoofs, scanner.Text())
	}
	if len(titleGoofs) == 0   {
		panic("Could not load any title goofs from " + filepath)
	}
}

// Return a randomly picked goof from the loaded title goofs
func GetRandomTitleGoof() string {
	return titleGoofs[rand.Intn(len(titleGoofs))]
}
