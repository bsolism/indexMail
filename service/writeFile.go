package service

import (
	"log"
	"os"
	"strings"
)

func WriteFile(errorResult string, path string, count int, nameFile string) {
	path = strings.ReplaceAll(string(path), "/", "_")
	f, err := os.Create("./error/" + path + nameFile + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err2 := f.WriteString(errorResult)
	if err2 != nil {
		log.Fatal(err2)
	}
}
