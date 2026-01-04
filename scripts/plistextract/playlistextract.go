package plistextract

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

const defaultDir = "C:\\Users\\throw\\Music\\Playlists"

// Exported function (uppercase) to be called outside the package
func ExtractPlaylists() ([]string, error) {
	dir := defaultDir

	x, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	//To store filenames
	var fileNames []string

	for _, entry := range x {
		if entry.IsDir() {
			continue //Skip directories
		}

		fileNames = append(fileNames, entry.Name())
	}

	return fileNames, nil
}

func ExtractSongs(filename string) ([]string, error) {
	dir := defaultDir
	fileDir := filepath.Join(dir, filename)

	file, err := os.Open(fileDir)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var songs []string

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "D:") {
			songs = append(songs, line)
		}
	}

	//Note to self - Learn how to handle scanner errors
	err = scanner.Err()
	if err != nil {
		return nil, err
	}
	return songs, nil
}
