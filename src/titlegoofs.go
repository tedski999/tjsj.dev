package main

import (
	"bufio"
	"os"
	"math/rand"
	"time"
)

var lines []string

func LoadTitleGoofs(filepath string) error {
	rand.Seed(time.Now().Unix())

	file, err := os.Open(filepath)
	if err != nil {
		return err
	}

	lines = nil
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	file.Close()
	return scanner.Err()
}

func GetRandomTitleGoof() string {
	return lines[rand.Intn(len(lines))]
}
