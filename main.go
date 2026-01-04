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

	//Extract songs
	songs, err := plistextract.ExtractSongs(selectedPlaylist)
	if err != nil {
		log.Fatalf("Error extracting songs: %v", err)
	}

	fmt.Println("Songs found:")
	for i, song := range songs {
		fmt.Printf("%d : %s\n", i+1, song)
	}
}
