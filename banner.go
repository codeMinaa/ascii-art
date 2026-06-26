package main

import (
	"errors"
	"os"
	"strings"
)

func LoadBanner(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("empty file")
	}
	lines := strings.Split(string(data), "\n")
	if len(lines) < 855 {
		return nil, errors.New("invalid banner format")
	}
	return lines, nil
}
