package content

import (
	"bufio"
	"os"
	"math/rand"
	"time"
)

var titleGoofs []string

// Load each line as a separate goof
func LoadTitleGoofs(filepath string) error {
	rand.Seed(time.Now().Unix())

	file, err := os.Open(filepath)
	if err != nil {
		return err
	}

	titleGoofs = nil
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		titleGoofs = append(titleGoofs, scanner.Text())
	}

	file.Close()
	return scanner.Err()
}

// Return a randomly picked goof from the loaded title goofs
func GetRandomTitleGoof() string {
	return titleGoofs[rand.Intn(len(titleGoofs))]
}
