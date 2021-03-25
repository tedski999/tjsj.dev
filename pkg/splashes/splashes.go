package splashes

import (
	"bufio"
	"os"
	"math/rand"
)

type Splashes struct {
	splashes []string
	random rand.Rand
}

// Load each line from filepath as a separate splash
func Load(filepath string) *Splashes {

	// Open file
	file, err := os.Open(filepath)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	// Read splashes from file
	s := []string {}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s = append(s, scanner.Text())
	}
	if len(s) == 0 {
		panic("Unable to load any splashes from " + filepath)
	}

	return &Splashes {
		splashes: s,
		random: *rand.New(rand.NewSource(0)),
	}
}

// Return a randomly picked line from the loaded splashes
func (s *Splashes) GetRandom() string {
	i := s.random.Intn(len(s.splashes))
	return s.splashes[i]
}
