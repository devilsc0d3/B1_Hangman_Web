module web

go 1.19

replace classic => ./hangman_classic_base

require (
	classic v0.0.0-00010101000000-000000000000
	github.com/hajimehoshi/go-mp3 v0.3.4
	github.com/hajimehoshi/oto v1.0.1
)

require (
	golang.org/x/exp v0.0.0-20190306152737-a1d7652674e8 // indirect
	golang.org/x/image v0.0.0-20190227222117-0694c2d4d067 // indirect
	golang.org/x/mobile v0.0.0-20190415191353-3e0bab5405d6 // indirect
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
)
