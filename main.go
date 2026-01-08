package main

import (
	"fmt"
	"log"

	"github.com/xcCodebreaker/rbsg/scripts/plistextract"
)

func main() {
	// Extract the playlists
	playlists, err := plistextract.ExtractPlaylists()
	if err != nil {
		log.Fatalf("Error extracting playlists: %v", err)
	}

	fmt.Println("Playlists found:")
	for i, playlist := range playlists {
		fmt.Printf("%d : %s\n", i+1, playlist)
	}

	//Select a playlist
	var choice int

	fmt.Println("Select a playlist (1, 2 or...): ")
	fmt.Scan(&choice)
	if choice < 1 || choice > len(playlists) {
		log.Fatal("Invalid playlist selection")
	}

	selectedPlaylist := playlists[choice-1]

	fmt.Println("Selected playlist:", selectedPlaylist)
	fmt.Println("")

	//Extract songs directory
	songs, err := plistextract.ExtractSongs(selectedPlaylist)
	if err != nil {
		log.Fatalf("Error extracting songs: %v", err)
	}

	var temp []string

	fmt.Println("Songs found:")
	for _, song := range songs {
		fmt.Printf("%s\n", song)
		temp = append(temp, song)
	}

	//Get the target directory
	var targetDir string

	fmt.Println("Target directory to copy the songs: ")
	fmt.Scan(&targetDir)

	//Copy the songs to another directory

	copied, err := plistextract.CopySongs(temp, targetDir)
	if err != nil {
		log.Fatalf("Error copying songs: %v", err) //Try not to forget error handling
	}

	fmt.Println("Copied files:")
	for _, c := range copied {
		fmt.Println(c)
	}
}
