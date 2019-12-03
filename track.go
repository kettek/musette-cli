package main

// Track represents, amazingly, a music track. This structure is used by the playerListView and is populated by our browser model.
type Track struct {
	trackNumber int8
	trackName   string
	trackArtist string
	trackAlbum  string
	selected    bool   // Whether or not this track is selected for operations
	path        string // Path to music file on the current musette server
}
