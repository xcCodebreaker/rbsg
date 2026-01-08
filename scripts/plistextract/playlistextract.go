package plistextract

import (
	"bufio"
	"io"
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

func CopySongs(songs []string, target string) ([]string, error) {
	targetDir := filepath.Clean(target)

	//Check if dir exists
	err := os.MkdirAll(targetDir, 0755)
	if err != nil {
		return nil, err
	}

	var copied []string

	for _, songsDir := range songs {
		srcPath := filepath.Clean(songsDir)

		//Open the file/song
		src, err := os.Open(srcPath)
		if err != nil {
			return nil, err
		}

		//Create the destination path - this part is a bit hazy but I think I get it. Need to dive deeper.
		//Write the destination path first
		dstPath := filepath.Join(targetDir, filepath.Base(srcPath))

		//Now, create the destination path
		dst, err := os.Create(dstPath)
		if err != nil {
			src.Close() //Why this? I have to start writing the doc soon
			return nil, err
		}

		//Copy bytes
		_, err = io.Copy(dst, src)

		//Close files
		src.Close()
		dst.Close() //Why not use defer here? where does the function end?

		if err != nil {
			return nil, err
		}

		copied = append(copied, dstPath)
	}

	return copied, nil
}
