package utils

import (
	"log"
	"os"
	"strings"
)

func TransformFile(file string) []string {
	content, err := os.ReadFile(file)

	if err != nil {
		log.Fatal(err)
	}

	return strings.Fields(string(content))
}
