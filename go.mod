module Hangman-Web

go 1.19

require (
	Hangman-Classic v0.0.1
)

replace (
	Hangman-Classic => "../hangman-classic"
)